package tweens
import (
	"time"
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

func MoveTo(movable Movable, x float64, y float64, duration time.Duration) *MoveToCmd {
	startX, startY := movable.GetPosition()
	return &MoveToCmd{subject: movable, funX: FromTo(startX, x, duration), funY: FromTo(startY, y, duration)}
}

func FromTo(from float64, to float64, duration time.Duration) func(tick time.Duration) float64 {

	return func(tick time.Duration) float64 {
		if tick <= 0 {
			return from
		}
		if tick >= duration {
			return to
		}
		return from + float64(tick) / float64(duration) * to
	}
}