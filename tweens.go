package tweens
import (
	"math"
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
	SetPosition(x int, y int)
	GetPosition() (x int, y int)
}

type MoveToCmd struct {
	subject Movable
	funX    func(time.Duration) int
	funY    func(time.Duration) int
}

func (m *MoveToCmd) Set(tick time.Duration) {
	m.subject.SetPosition(m.funX(tick), m.funY(tick))
}

func MoveTo(movable Movable, x int, y int, duration time.Duration) *MoveToCmd {
	startX, startY := movable.GetPosition()
	return &MoveToCmd{subject: movable, funX: FromTo(startX, x, duration), funY: FromTo(startY, y, duration)}
}

// TODO: make from and to float64
func FromTo(from int, to int, duration time.Duration) func(tick time.Duration) int {

	return func(tick time.Duration) int {
		if tick <= 0 {
			return from
		}
		if tick > duration {
			return to
		}
		result := round2int(float64(from) + float64(tick) / float64(duration) * float64(to))
		return result
	}
}

func round2int(input float64) int {
	if input < 0 {
		return int(math.Ceil(input - 0.5))
	}
	return int(math.Floor(input + 0.5))
}