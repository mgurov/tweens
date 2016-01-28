package tweens

import (
	"time"
)

// Changes progress over the time from the moment zero which is normall when the Scene has started
type Change interface {
	Progress(tick time.Duration)
}

// the terminal goal of the change, e.g. move from A to B
// maps completion [0,1] -> [A,B]
type TransitionCompletionFunction func(completion float64)

func NoopCompletion(completion float64) {}

type Transition interface {
	// Transition might depends on the initial state of the subject
	// which in turn might depend on the time, especially in sequences
	// therefore this method is called only when the transition begins
	Start() TransitionCompletionFunction
}

type TickNormalizationFunction func(tick time.Duration) float64

type How struct {
	Duration   time.Duration
	Easing     Easing
	Repetition Lifespan
}
