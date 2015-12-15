package tweens
import (
	"time"
	"math"
)

type Scene struct {
	items []Setter
}

func (s *Scene) Add(newSetter Setter) {
	s.items = append(s.items, newSetter)
}

// Sets the timestamp
func (s *Scene) Set(t time.Duration) {
	for _, m := range s.items {
		m.Set(t)
	}
}

type Setter interface {
	Set(tick time.Duration)
}

type Movable interface {
	SetPosition(x float64, y float64)
	GetPosition() (x float64, y float64)
}

type MoveToCmd struct {
	subject Movable
	funX    func(time.Duration) float64
	funY    func(time.Duration) float64
}

func (m *MoveToCmd) Set(tick time.Duration) {
	m.subject.SetPosition(m.funX(tick), m.funY(tick))
}

func MoveTo(movable Movable, x float64, y float64, duration time.Duration, easing Easing) *MoveToCmd {
	return MoveToRepeat(movable, x, y, duration, easing, Once)
}

func MoveToRepeat(movable Movable, x float64, y float64, duration time.Duration, easing Easing, repeat Lifespan) *MoveToCmd {
	startX, startY := movable.GetPosition()
	return &MoveToCmd{subject: movable, funX: FromTo(startX, x, duration, easing, repeat), funY: FromTo(startY, y, duration, easing, repeat)}
}

func FromTo(from float64, to float64, duration time.Duration, easing Easing, lifespan Lifespan) func(tick time.Duration) float64 {

	return func(tick time.Duration) float64 {
		if tick <= 0 {
			return from
		}
		completed := float64(tick) / float64(duration)

		return from + easing(lifespan(completed)) * to
	}
}

type Lifespan func (in float64) float64

func Once (in float64) float64 {
	if in > 1 {
		return 1
	} else {
		return in
	}
}

func Repeat (in float64) float64 {
	_,fraction := math.Modf(in)
	if 0 == fraction {
		return 1
	} else {
		return fraction
	}
}

func YoYo (in float64) float64 {
	whole,fraction := math.Modf(in)
	if 0 == (int)(whole) % 2 {
		return fraction
	} else {
		return 1 - fraction
	}
}