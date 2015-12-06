package tweens

import (
	"fmt"
)

type Sprite struct {
	x int
	y int
}

func (s *Sprite) SetPosition(x int, y int) {
	s.x = x
	s.y = y
}

func (s *Sprite) GetPosition() (x int, y int) {
	return s.x, s.y
}

func execute(s *Scene, toTick int, callback func(tick int)) {
	for t := 0; t <= toTick; t++ {
		s.Set(t)
		callback(t)
	}
}

func ExampleSimplest() {

	s := Sprite{0, 0}

	scene := Scene{}

	scene.Add(MoveTo(&s, 100, -100, 5))

	execute(&scene, 6, func(tick int) {
		fmt.Println(s.x, s.y)
	})

	// Output:
	//0 0
	//20 -20
	//40 -40
	//60 -60
	//80 -80
	//100 -100
	//100 -100
}

func ExampleTinyStep() {

	s := Sprite{0, 0}

	scene := Scene{}

	scene.Add(MoveTo(&s, 100, 1, 4))

	execute(&scene, 6, func(tick int) {
		fmt.Println(s.x, s.y)
	})

	// Output:
	//0 0
	//25 0
	//50 1
	//75 1
	//100 1
	//100 1
	//100 1
}