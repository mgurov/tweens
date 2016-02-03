package tweens

import (
	"time"
)

type Sequence struct {
	Steps            []*Step
	Repetition       RepetitionFunction
	currentStepIndex int
	duration         time.Duration
}

func NewSequence(steps ...*Step) *Sequence {
	s := &Sequence{
		Steps:            steps,
		currentStepIndex: -1,
		Repetition:       Once,
		duration:         prepareSteps(steps),
	}

	return s
}

func (s *Sequence) YoYo() *Sequence {

	s.Repetition = YoYo

	return s
}

func (s *Sequence) Repeat() *Sequence {

	s.Repetition = Repeat

	return s
}

func (s *Sequence) Progress(tick time.Duration) {

	if tick < 0 {
		return
	}

	if tick > s.duration && s.duration > 0 {
		tick = s.Repetition(tick, s.duration)
	}

	step := s.findStep(tick)
	if step == nil {
		return
	}

	stepTick := tick - step.start
	var completed float64
	if stepTick <= 0 {
		completed = 0
	} else {
		completed = float64(stepTick) / float64(step.Duration)
	}

	step.apply(completed)
}

func (s *Sequence) findStep(tick time.Duration) *Step {

	beforeStart := s.currentStepIndex < 0
	afterEnd := s.currentStepIndex >= len(s.Steps)

	var ascending bool

	if beforeStart || afterEnd {
		ascending = beforeStart || tick > s.duration
	} else {
		slot := s.Steps[s.currentStepIndex]
		if slot.covers(tick) {
			return slot
		}
		ascending = slot.stop < tick
	}

	cur := cursor{Sequence: s, ascending: ascending}
	return cur.findSlot(tick)

}

type cursor struct {
	*Sequence
	ascending bool
}

func (c *cursor) isEndOfSlots() bool {
	return c.currentStepIndex >= len(c.Steps) || c.currentStepIndex < 0
}

func (c *cursor) moveNextSlot() {
	if c.ascending {
		c.currentStepIndex++
	} else {
		c.currentStepIndex--
	}
}

func (c *cursor) currentSlot() *Step {
	if c.isEndOfSlots() {
		return nil
	} else {
		return c.Steps[c.currentStepIndex]
	}
}

func (c *cursor) findSlot(tick time.Duration) *Step {

	var completionValueForPassedTicks float64
	if c.ascending {
		completionValueForPassedTicks = fullyComplete
	} else {
		completionValueForPassedTicks = atStart
	}

	for {
		step := c.currentSlot()
		if nil != step {
			if step.covers(tick) {
				return step
			} else {
				step.apply(completionValueForPassedTicks)
			}
		}
		c.moveNextSlot()
		if c.isEndOfSlots() {
			return nil
		}
	}
}

const fullyComplete float64 = 1.0
const atStart float64 = 0.0
