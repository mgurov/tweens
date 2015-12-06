package tweens
import (
	"math"
)

type Scene struct {
	items []Setter
}

func (s *Scene) Add(newSetter Setter) {
	s.items = append(s.items, newSetter)
}

// Sets the timestamp
func (s *Scene) Set(t int) {
	for _, m := range s.items {
		m.Set(t)
	}
}

type Setter interface {
	Set(tick int)
}

type Movable interface {
	SetPosition(x int, y int)
	GetPosition() (x int, y int)
}

type MoveToCmd struct {
	subject Movable
	funX    func(int) int
	funY    func(int) int
}

func (m *MoveToCmd) Set(tick int) {
	m.subject.SetPosition(m.funX(tick), m.funY(tick))
}

func MoveTo(movable Movable, x int, y int, duration int) *MoveToCmd {
	startX, startY := movable.GetPosition()
	return &MoveToCmd{subject: movable, funX: FromTo(startX, x, duration), funY: FromTo(startY, y, duration)}
}

func FromTo(from int, to int, duration int) func(tick int) int {

	delta := float64(to - from) / float64(duration) //TODO: division by zero

	return func(tick int) int {
		if tick > duration {
			return to
		}
		result := round2int(float64(from) + delta * float64(tick))
		return result
	}
}

func round2int(input float64) int {
	if input < 0 {
		return int(math.Ceil(input - 0.5))
	}
	return int(math.Floor(input + 0.5))
}