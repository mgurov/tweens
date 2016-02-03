package tweens

import (
	"time"
)

func Pause(duration time.Duration) *Step {
	return &Step{What: pause{}, Duration: duration}
}

type pause struct {
}

func (_ pause) Start() TransitionCompletionFunction {
	return NoopCompletion
}
