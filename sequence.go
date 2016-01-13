package tweens
import (
	"time"
	"math"
)

type Sequence struct {
	Steps [] Step
}

type Step struct {
	What     func() ChangeFunction
	Duration time.Duration
	Easing   Easing
}

type timeLine struct {
	timeSlots        []*timeSlot
	currentStepIndex int
	lifespan         Lifespan
	duration         time.Duration
}

type timeSlot struct {
	start        time.Duration
	stop         time.Duration
	duration     time.Duration
	whatProvider func() ChangeFunction
	what         ChangeFunction
	easing       Easing
}

//TODO: naming
func (s *timeSlot) getOrMakeWhat() ChangeFunction {
	if s.what == nil {
		s.what = s.whatProvider()
	}
	return s.what
}

func (s *timeLine) Set(tick time.Duration) {

	//TODO: make sane
	if tick > s.duration {
		ratio := s.lifespan(float64(tick) / float64(s.duration))
		tick = time.Duration(float64(s.duration) * s.lifespan(ratio))
	}

	slot := s.findChangeFunction(tick)
	if slot == nil {
		return
	}

	stepTick := tick - slot.start
	var completed float64
	if stepTick <= 0 {
		//shift the tiniest bit to make downstream functions happier. Kind of a special case, although might need rethinking.
		completed = math.Nextafter(0, 1)
	} else {
		completed = float64(stepTick) / float64(slot.duration)
	}

	slot.getOrMakeWhat()(slot.easing(completed))
}

func (s *timeLine) findChangeFunction(tick time.Duration) *timeSlot {
	for i, slot := range s.timeSlots {
		if slot.start <= tick && slot.stop >= tick {
			//first complete the continuum for all the previous steps just in case we weren't thorough enough with updating them
			//TODO: ranges maybe?
			if i > s.currentStepIndex {
				if s.currentStepIndex < 0 {
					s.currentStepIndex = 0
				}
				for r := s.currentStepIndex; r < i; r++ {
					s.timeSlots[r].getOrMakeWhat()(1)
				}
			} else if i < s.currentStepIndex {
				if s.currentStepIndex >= len(s.timeSlots) {
					s.currentStepIndex = len(s.timeSlots) - 1
				}
				for r := i + 1; r <= s.currentStepIndex; r++ {
					s.timeSlots[r].getOrMakeWhat()(0)
				}
			}
			s.currentStepIndex = i
			return slot
		}
	}
	if tick > 0 {
		s.currentStepIndex = len(s.timeSlots) + 1
	} else {
		s.currentStepIndex = -1
	}
	return nil
}

func (s *Sequence) Build(lifespan Lifespan) Setter {

	if len(s.Steps) == 0 {
		//NOOP
		return setter{fun: func(tick time.Duration) {}}
	}

	//TODO: optimize
	var timeCursor time.Duration
	timeSlots := make([]*timeSlot, len(s.Steps))
	var totalDuration time.Duration
	for i, step := range s.Steps {
		//TODO: squashing on duration 0?
		thisTimeSlot := timeSlot{
			start: timeCursor,
			stop: timeCursor + step.Duration,
			duration: step.Duration,
			whatProvider: step.What,
			easing: step.Easing,
		}

		if nil == thisTimeSlot.easing {
			thisTimeSlot.easing = Linear
		}

		timeCursor = thisTimeSlot.stop //TODO: maybe add minimal instead of simply starting from the same point?
		totalDuration += thisTimeSlot.duration
		timeSlots[i] = &thisTimeSlot
	}

	return &timeLine{timeSlots: timeSlots, currentStepIndex: -1, lifespan: lifespan, duration: totalDuration}
}

func (s *Sequence) Once() Setter {
	return s.Build(Once)
}

type setter struct {
	fun func(tick time.Duration)
}

func (s setter) Set(tick time.Duration) {
	s.fun(tick)
}