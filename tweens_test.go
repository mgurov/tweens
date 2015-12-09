package tweens

import (
	"fmt"
	"time"
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

func execute(s *Scene, toSecond int, callback func(tick int)) {
	for t := 0; t <= toSecond; t++ {
		s.Set(time.Duration(t) * time.Second)
		callback(t)
	}
}

func ExampleSimplest() {

	s := Sprite{0, 0}

	scene := Scene{}

	scene.Add(MoveTo(&s, 100, -100, time.Duration(5) * time.Second))

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

	scene.Add(MoveTo(&s, 100, 1, time.Duration(4) * time.Second))

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

func ExampleNegativeStepsIgnored() {

	s := Sprite{0, 0}

	scene := Scene{}

	scene.Add(MoveTo(&s, 100, 1, time.Duration(4) * time.Second))

	scene.Set(time.Duration(-1) * time.Second)

	fmt.Println(s.x, s.y)

	// Output:
	//0 0
}

func ExampleZeroDuration() {

	s := Sprite{0, 0}

	scene := Scene{}

	scene.Add(MoveTo(&s, 100, 1, time.Duration(0)))

	execute(&scene, 2, func(tick int) {
		fmt.Println(s.x, s.y)
	})

	// Output:
	//0 0
	//100 1
	//100 1
}
