package tweens

type TweenAspect int

// Tweenable is a convenience interface that allows covering several "tweenable" aspects with a single implementation method
type Tweenable interface {
	GetValues(aspect TweenAspect) []float64
	SetValues(aspect TweenAspect, newValues []float64)
}

// Predefined tweenable aspects, for demonstration and convenience purpose
// One can use own constants mixing or overlapping with these
const (
	AspectPosition TweenAspect = iota
	AspectSize
	AspectColor
	AspectDirection
)

func Move(subject Tweenable, x, y float64) Transition {
	return NewTweenableAccessor(subject, AspectPosition, x, y)
}

func Resize(subject Tweenable, h, w float64) Transition {
	return NewTweenableAccessor(subject, AspectSize, h, w)
}

func Rotate(subject Tweenable, angle float64) Transition {
	return NewTweenableAccessor(subject, AspectDirection, angle)
}

func Colorize(subject Tweenable, r, g, b int) Transition {
	return NewTweenableAccessor(subject, AspectColor, float64(r), float64(g), float64(b))
}

type tweenableAccessor struct {
	subject Tweenable
	aspect  TweenAspect
}

func NewTweenableAccessor(subject Tweenable, aspect TweenAspect, targets ...float64) Transition {
	return LazyAccessor(&tweenableAccessor{subject, aspect}, targets...)
}

func (t *tweenableAccessor) Get() (currentState []float64) {
	return t.subject.GetValues(t.aspect)
}

func (t *tweenableAccessor) Set(newState []float64) {
	t.subject.SetValues(t.aspect, newState)
}
