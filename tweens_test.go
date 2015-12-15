package tweens

import (
	"fmt"
	"time"
	"math"
)

func ExampleSimplest() {

	s := Sprite{0, 0}

	scene := Scene{}

	scene.Add(MoveTo(&s, 100, -100, time.Duration(5) * time.Second, Linear))

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

	scene.Add(MoveTo(&s, 100, 1, time.Duration(4) * time.Second, Linear))

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

	scene.Add(MoveTo(&s, 100, 1, time.Duration(4) * time.Second, Linear))

	scene.Set(time.Duration(-1) * time.Second)

	fmt.Println(s.x, s.y)

	// Output:
	//0 0
}

func ExampleZeroDuration() {

	s := Sprite{0, 0}

	scene := Scene{}

	scene.Add(MoveTo(&s, 100, 1, time.Duration(0), Linear))

	execute(&scene, 2, func(tick int) {
		fmt.Println(s.x, s.y)
	})

	// Output:
	//0 0
	//100 1
	//100 1
}

func ExampleEaseInQuad() {

	s := Sprite{0, 0}

	scene := Scene{}

	scene.Add(MoveTo(&s, 100, -100, time.Duration(5) * time.Second, EaseInQuad))

	execute(&scene, 6, func(tick int) {
		fmt.Println(s.x, s.y)
	})

	// Output:
	//0 0
	//4 -4
	//16 -16
	//36 -36
	//64 -64
	//100 -100
	//100 -100
}

func ExampleRepeat() {

	s := Sprite{0, 0}

	scene := Scene{}

	scene.Add(MoveToRepeat(&s, 100, -100, time.Duration(5) * time.Second, Linear, Repeat))

	execute(&scene, 11, func(tick int) {
		fmt.Println(s.x, s.y)
	})

	// Output:
	//0 0
	//20 -20
	//40 -40
	//60 -60
	//80 -80
	//100 -100
	//20 -20
	//40 -40
	//60 -60
	//80 -80
	//100 -100
	//20 -20
}

func ExampleYoYo() {

	s := Sprite{0, 0}

	scene := Scene{}

	scene.Add(MoveToRepeat(&s, 100, -100, time.Duration(5) * time.Second, Linear, YoYo))

	execute(&scene, 11, func(tick int) {
		fmt.Println(s.x, s.y)
	})

	// Output:
	//0 0
	//20 -20
	//40 -40
	//60 -60
	//80 -80
	//100 -100
	//80 -80
	//60 -60
	//40 -40
	//20 -20
	//0 0
	//20 -20
}

func execute(s *Scene, toSecond int, callback func(tick int)) {
	for t := 0; t <= toSecond; t++ {
		s.Set(time.Duration(t) * time.Second)
		callback(t)
	}
}

type Sprite struct {
	x int
	y int
}

func (s *Sprite) SetPosition(x float64, y float64) {
	s.x = round2int(x)
	s.y = round2int(y)
}

func (s *Sprite) GetPosition() (x float64, y float64) {
	return float64(s.x), float64(s.y)
}

func round2int(input float64) int {
	if input < 0 {
		return int(math.Ceil(input - 0.5))
	}
	return int(math.Floor(input + 0.5))
}
