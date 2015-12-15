package tweens
import (
	"math"
)

// adopted from https://github.com/gopackage/tween/blob/master/curves/ease.go
type Easing func(completed float64) float64

func Linear(completed float64) float64 {
	return completed
}

func EaseInQuad(completed float64) float64 {
	return math.Pow(completed, 2)
}


// EaseInOutBounce eases in and out a Bounce transition.
// See http://jqueryui.com/easing/ for curve in action.
func EaseInOutBounce(completed float64) float64 {
	if completed < 0.5 {
		return EaseInBounce(completed*2) / 2
	}
	return 1 - EaseInBounce((completed*-2)+2)/2
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