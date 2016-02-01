package tweens

import (
	"time"
)

//of a sequence
type Step struct {
	What     Transition
	Duration time.Duration
	Easing   Easing
	start    time.Duration
	stop     time.Duration
	what     TransitionCompletionFunction //TODO: this way we cache the transition, but what if I want to do new every time? Stacking or something?
}

// sets intervals, defaults and counts the total duration of the sequence
func prepareSteps(steps []*Step) (totalDuration time.Duration) {
	var timeCursor time.Duration
	for _, step := range steps {
		step.start = timeCursor
		step.stop = timeCursor + step.Duration //TODO: should be bit more precise on whether the stop includes or not.
		if nil == step.Easing {
			step.Easing = Linear
		}
		timeCursor = step.stop
		totalDuration += step.Duration
	}
	return
}

func (t Step) covers(tick time.Duration) bool {
	return t.start <= tick && t.stop >= tick
}

func (s *Step) apply(complete float64) {
	if s.what == nil {
		s.what = s.What.Start()
	}

	s.what(s.Easing(complete))
}