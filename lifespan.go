package tweens
import "math"

// Lifespan defines the behavior of the transformation beyond the base duration, e.g. repeat, yoyo.
type Lifespan func(in float64) float64

func Once(in float64) float64 {
	if in > 1 {
		return 1
	} else {
		return in
	}
}

func Repeat(in float64) float64 {
	_, fraction := math.Modf(in)
	if 0 == fraction {
		return 1
	} else {
		return fraction
	}
}

func YoYo(in float64) float64 {
	whole, fraction := math.Modf(in)
	if 0 == (int)(whole) % 2 {
		return fraction
	} else {
		return 1 - fraction
	}
}