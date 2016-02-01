package tweens

import (
	"time"
)

// Lifespan defines the behavior of the transformation beyond the base duration, e.g. repeat, yoyo.
type Lifespan func(now time.Duration, span time.Duration) time.Duration

func Once(now time.Duration, span time.Duration) time.Duration {
	if now > span {
		return span
	} else {
		return now
	}
}

func Repeat(now time.Duration, span time.Duration) time.Duration {
	reminder := now % span
	if 0 == reminder {
		return span
	} else {
		return reminder
	}
}

func YoYo(now time.Duration, span time.Duration) time.Duration {
	reminder := now % span
	whole := now / span
	if 0 == whole % 2 {
		return reminder
	} else {
		return span - reminder
	}
}
