package tweens

import (
	"math"
	"time"
)

type Sequence struct {
	Steps []Step
}

func NewSequence(steps ...Step) *Sequence {
	return &Sequence{
		Steps: steps,
	}
}

type Step struct {
	What     Transition
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
	start    time.Duration
	stop     time.Duration
	duration time.Duration
	change   Transition
	what     TransitionCompletionFunction //TODO: this way we cache the transition, but what if I want to do new every time? Stacking or something?
	easing   Easing
}

func (t *timeSlot) covers(tick time.Duration) bool {
	return t.start <= tick && t.stop >= tick
}

//TODO: naming
func (s *timeSlot) getOrMakeWhat() TransitionCompletionFunction {
	if s.what == nil {
		s.what = s.change.Start()
	}
	return s.what
}

func (s *timeLine) Progress(tick time.Duration) {

	if tick < 0 {
		return
	}

	//TODO: make sane
	if tick > s.duration && s.duration > 0 {
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

	beforeStart := s.currentStepIndex < 0
	afterEnd := s.currentStepIndex > len(s.timeSlots)

	var ascending bool

	if beforeStart || afterEnd {
		ascending = beforeStart || tick > s.duration
	} else {
		slot := s.timeSlots[s.currentStepIndex]
		if slot.covers(tick) {
			return slot
		}
		ascending = slot.stop < tick
	}

	if ascending {
		if beforeStart {
			s.currentStepIndex = 0
		}
		for ; s.currentStepIndex < len(s.timeSlots); s.currentStepIndex++ {
			slot := s.timeSlots[s.currentStepIndex]
			if slot.covers(tick) {
				return slot
			} else {
				slot.getOrMakeWhat()(1)
			}
		}
		s.currentStepIndex++
	} else {
		if afterEnd {
			s.currentStepIndex = len(s.timeSlots) - 1
		}
		for ; s.currentStepIndex >= 0; s.currentStepIndex-- {
			slot := s.timeSlots[s.currentStepIndex]
			if slot.covers(tick) {
				return slot
			} else {
				slot.getOrMakeWhat()(0)
			}
		}
		s.currentStepIndex--
	}

	return nil
}

func (s *Sequence) Build(lifespan Lifespan) Change {

	if len(s.Steps) == 0 {
		//NOOP
		return setter{fun: func(tick time.Duration) {}}
	}

	if nil == lifespan {
		lifespan = Linear
	}

	//TODO: optimize
	var timeCursor time.Duration
	timeSlots := make([]*timeSlot, len(s.Steps))
	var totalDuration time.Duration
	for i, step := range s.Steps {
		//TODO: squashing on duration 0?
		thisTimeSlot := timeSlot{
			start:    timeCursor,
			stop:     timeCursor + step.Duration,
			duration: step.Duration,
			change:   step.What,
			easing:   step.Easing,
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

func (s *Sequence) Once() Change {
	return s.Build(Once)
}

//TODO: move as a helper to the tweens
type setter struct {
	fun func(tick time.Duration)
}

func (s setter) Progress(tick time.Duration) {
	s.fun(tick)
}

func Pause(duration time.Duration) Step {
	//TODO: or maybe do What by default nothing?
	return Step{What: pause{}, Duration: duration}
}

type pause struct {
}

func (_ pause) Start() TransitionCompletionFunction {
	return NoopCompletion
}
