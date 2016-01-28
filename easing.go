package tweens

import (
	"math"
)

type Easing func(completed float64) float64

func Linear(completed float64) float64 {
	return completed
}

// below adaptation by copying from https://github.com/gopackage/tween/blob/master/curves/ease.go

// EaseInQuad eases in a Quad transition.
// See http://jqueryui.com/easing/ for curve in action.
func EaseInQuad(completed float64) float64 {
	return math.Pow(completed, 2)
}

// EaseOutQuad eases out a Quad transition.
// See http://jqueryui.com/easing/ for curve in action.
func EaseOutQuad(completed float64) float64 {
	return 1 - EaseInQuad(1 - completed)
}

// EaseInOutQuad eases in and out a Quad transition.
// See http://jqueryui.com/easing/ for curve in action.
func EaseInOutQuad(completed float64) float64 {
	if completed < 0.5 {
		return EaseInQuad(completed * 2) / 2
	}
	return 1 - EaseInQuad((completed * -2) + 2) / 2
}

// EaseInCubic eases in a Cubic transition.
// See http://jqueryui.com/easing/ for curve in action.
func EaseInCubic(completed float64) float64 {
	return math.Pow(completed, 3)
}

// EaseOutCubic eases out a Cubic transition.
// See http://jqueryui.com/easing/ for curve in action.
func EaseOutCubic(completed float64) float64 {
	return 1 - EaseInCubic(1 - completed)
}

// EaseInOutCubic eases in and out a Cubic transition.
// See http://jqueryui.com/easing/ for curve in action.
func EaseInOutCubic(completed float64) float64 {
	if completed < 0.5 {
		return EaseInCubic(completed * 2) / 2
	}
	return 1 - EaseInCubic((completed * -2) + 2) / 2
}

// EaseInQuart eases in a Quart transition.
// See http://jqueryui.com/easing/ for curve in action.
func EaseInQuart(completed float64) float64 {
	return math.Pow(completed, 4)
}

// EaseOutQuart eases out a Quart transition.
// See http://jqueryui.com/easing/ for curve in action.
func EaseOutQuart(completed float64) float64 {
	return 1 - EaseInQuart(1 - completed)
}

// EaseInOutQuart eases in and out a Quart transition.
// See http://jqueryui.com/easing/ for curve in action.
func EaseInOutQuart(completed float64) float64 {
	if completed < 0.5 {
		return EaseInQuart(completed * 2) / 2
	}
	return 1 - EaseInQuart((completed * -2) + 2) / 2
}

// EaseInQuint eases in a Quint transition.
// See http://jqueryui.com/easing/ for curve in action.
func EaseInQuint(completed float64) float64 {
	return math.Pow(completed, 5)
}

// EaseOutQuint eases out a Quint transition.
// See http://jqueryui.com/easing/ for curve in action.
func EaseOutQuint(completed float64) float64 {
	return 1 - EaseInQuint(1 - completed)
}

// EaseInOutQuint eases in and out a Quint transition.
// See http://jqueryui.com/easing/ for curve in action.
func EaseInOutQuint(completed float64) float64 {
	if completed < 0.5 {
		return EaseInQuint(completed * 2) / 2
	}
	return 1 - EaseInQuint((completed * -2) + 2) / 2
}

// EaseInExpo eases in a Expo transition.
// See http://jqueryui.com/easing/ for curve in action.
func EaseInExpo(completed float64) float64 {
	return math.Pow(completed, 6)
}

// EaseOutExpo eases out a Expo transition.
// See http://jqueryui.com/easing/ for curve in action.
func EaseOutExpo(completed float64) float64 {
	return 1 - EaseInExpo(1 - completed)
}

// EaseInOutExpo eases in and out a Expo transition.
// See http://jqueryui.com/easing/ for curve in action.
func EaseInOutExpo(completed float64) float64 {
	if completed < 0.5 {
		return EaseInExpo(completed * 2) / 2
	}
	return 1 - EaseInExpo((completed * -2) + 2) / 2
}

// EaseInSine eases in a Sine transition.
// See http://jqueryui.com/easing/ for curve in action.
func EaseInSine(completed float64) float64 {
	return 1 - math.Cos(completed * math.Pi / 2)
}

// EaseOutSine eases out a Sine transition.
// See http://jqueryui.com/easing/ for curve in action.
func EaseOutSine(completed float64) float64 {
	return 1 - EaseInSine(1 - completed)
}

// EaseInOutSine eases in and out a Sine transition.
// See http://jqueryui.com/easing/ for curve in action.
func EaseInOutSine(completed float64) float64 {
	if completed < 0.5 {
		return EaseInSine(completed * 2) / 2
	}
	return 1 - EaseInSine((completed * -2) + 2) / 2
}

// EaseInCirc eases in a Circ transition.
// See http://jqueryui.com/easing/ for curve in action.
func EaseInCirc(completed float64) float64 {
	return 1 - math.Sqrt(1 - completed * completed)
}

// EaseOutCirc eases out a Circ transition.
// See http://jqueryui.com/easing/ for curve in action.
func EaseOutCirc(completed float64) float64 {
	return 1 - EaseInCirc(1 - completed)
}

// EaseInOutCirc eases in and out a Circ transition.
// See http://jqueryui.com/easing/ for curve in action.
func EaseInOutCirc(completed float64) float64 {
	if completed < 0.5 {
		return EaseInCirc(completed * 2) / 2
	}
	return 1 - EaseInCirc((completed * -2) + 2) / 2
}

// EaseInElastic eases in a Elastic transition.
// See http://jqueryui.com/easing/ for curve in action.
func EaseInElastic(completed float64) float64 {
	if completed == 0 || completed == 1 {
		return completed
	}
	return -math.Pow(2, 8 * (completed - 1)) * math.Sin(((completed - 1) * 80 - 7.5) * math.Pi / 15)
}

// EaseOutElastic eases out a Elastic transition.
// See http://jqueryui.com/easing/ for curve in action.
func EaseOutElastic(completed float64) float64 {
	return 1 - EaseInElastic(1 - completed)
}

// EaseInOutElastic eases in and out a Elastic transition.
// See http://jqueryui.com/easing/ for curve in action.
func EaseInOutElastic(completed float64) float64 {
	if completed < 0.5 {
		return EaseInElastic(completed * 2) / 2
	}
	return 1 - EaseInElastic((completed * -2) + 2) / 2
}

// EaseInBack eases in a Back transition.
// See http://jqueryui.com/easing/ for curve in action.
func EaseInBack(completed float64) float64 {
	return completed * completed * (3 * completed - 2)
}

// EaseOutBack eases out a Back transition.
// See http://jqueryui.com/easing/ for curve in action.
func EaseOutBack(completed float64) float64 {
	return 1 - EaseInBack(1 - completed)
}

// EaseInOutBack eases in and out a Back transition.
// See http://jqueryui.com/easing/ for curve in action.
func EaseInOutBack(completed float64) float64 {
	if completed < 0.5 {
		return EaseInBack(completed * 2) / 2
	}
	return 1 - EaseInBack((completed * -2) + 2) / 2
}

// EaseInBounce eases in a Bounce transition.
// See http://jqueryui.com/easing/ for curve in action.
func EaseInBounce(completed float64) float64 {

	bounce := float64(3)
	var pow2 float64
	for pow2 = math.Pow(2, bounce); completed < ((pow2 - 1) / 11); pow2 = math.Pow(2, bounce) {
		bounce--
	}
	return 1/math.Pow(4, 3-bounce) - 7.5625*math.Pow((pow2*3-2)/22-completed, 2)
}

// EaseOutBounce eases out a Bounce transition.
// See http://jqueryui.com/easing/ for curve in action.
func EaseOutBounce(completed float64) float64 {
	return 1 - EaseInBounce(1-completed)
}

// EaseInOutBounce eases in and out a Bounce transition.
// See http://jqueryui.com/easing/ for curve in action.
func EaseInOutBounce(completed float64) float64 {
	if completed < 0.5 {
		return EaseInBounce(completed * 2) / 2
	}
	return 1 - EaseInBounce((completed * -2) + 2) / 2
}
