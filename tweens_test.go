package tweens_test

import (
	"fmt"
	. "github.com/mgurov/tweens"
	"math"
	"testing"
	"time"
)

func ExampleSimplest() {

	s := Sprite{0, 0}

	scene := Scene{}

	scene.AddTransition(MoveTo2(&s, 100, -100), How{Duration: time.Duration(5) * time.Second})

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

	scene.AddTransition(MoveTo2(&s, 100, 1), How{Duration: time.Duration(4) * time.Second})

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

	scene.AddTransition(MoveTo2(&s, 100, 1), How{Duration: time.Duration(4) * time.Second})

	scene.Set(time.Duration(-1) * time.Second)

	fmt.Println(s.x, s.y)

	// Output:
	//0 0
}

func ExampleZeroDuration() {

	s := Sprite{0, 0}

	scene := Scene{}

	scene.AddTransition(MoveTo2(&s, 100, 1), How{Duration: time.Duration(0)})

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

	scene.AddTransition(MoveTo2(&s, 100, -100), How{Duration: time.Duration(5) * time.Second, Easing: EaseInQuad})

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

	scene.AddTransition(MoveTo2(&s, 100, -100), How{Duration: time.Duration(5) * time.Second, Repetition: Repeat})

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

	scene.AddTransition(MoveTo2(&s, 100, -100), How{Duration: time.Duration(5) * time.Second, Repetition: YoYo})

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

func ExampleSequence() {

	s := Sprite{0, 0}

	scene := Scene{}

	sequence := Sequence{
		[]Step{
			Step{What: MoveTo2(&s, 100, -100), Duration: time.Duration(5) * time.Second},
			Step{What: MoveTo2(&s, -100, 100), Duration: time.Duration(10) * time.Second},
		},
	}

	scene.Add(sequence.Once())

	execute(&scene, 17, func(tick int) {
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
	//-20 20
	//-40 40
	//-60 60
	//-80 80
	//-100 100
	//-100 100
	//-100 100
}

func ExampleRepeatSequence() {

	s := Sprite{0, 0}

	scene := Scene{}

	sequence := Sequence{
		[]Step{
			Step{What: MoveTo2(&s, 100, -100), Duration: 2 * time.Second},
			Step{What: MoveTo2(&s, -100, -200), Duration: 4 * time.Second},
		},
	}

	scene.Add(sequence.Build(Repeat))

	execute(&scene, 10, func(tick int) {
		fmt.Println(s.x, s.y)
	})

	// Output:
	//0 0
	//50 -50
	//100 -100
	//50 -125
	//0 -150
	//-50 -175
	//-100 -200
	//50 -50
	//100 -100
	//50 -125
	//0 -150
}

func ExampleYoYoSequence() {

	s := Sprite{0, 0}

	scene := Scene{}

	//TODO: repeat and such on the sequence.
	sequence := Sequence{
		[]Step{
			Step{What: MoveTo2(&s, 100, -100), Duration: 1 * time.Second},
			Step{What: MoveTo2(&s, -100, -200), Duration: 2 * time.Second},
		},
	}

	scene.Add(sequence.Build(YoYo))

	execute(&scene, 10, func(tick int) {
		fmt.Println(s.x, s.y)
	})

	// Output:
	//0 0
	//100 -100
	//0 -150
	//-100 -200
	//0 -150
	//100 -100
	//0 0
	//100 -100
	//0 -150
	//-100 -200
	//0 -150
}

func ExampleSequenceWithStepEasing() {

	s := Sprite{0, 0}

	scene := Scene{}

	//TODO: repeat and such on the sequence.
	sequence := Sequence{
		[]Step{
			Step{What: MoveTo2(&s, 100, -100), Duration: time.Duration(5) * time.Second, Easing: EaseInQuad},
			Step{What: MoveTo2(&s, -100, 100), Duration: time.Duration(10) * time.Second},
		},
	}

	scene.Add(sequence.Once())

	execute(&scene, 16, func(tick int) {
		fmt.Println(s.x, s.y)
	})

	// Output:
	//0 0
	//4 -4
	//16 -16
	//36 -36
	//64 -64
	//100 -100
	//80 -80
	//60 -60
	//40 -40
	//20 -20
	//0 0
	//-20 20
	//-40 40
	//-60 60
	//-80 80
	//-100 100
	//-100 100
}

func TestEmptySequence(t *testing.T) {

	scene := Scene{}

	setter := (&Sequence{}).Once()

	scene.Add(setter)

	scene.Set(1 * time.Second)
}

func TestCoarseSequenceShouldFastForwardIntermediateSteps(t *testing.T) {

	s1 := Sprite{0, 0}
	s2 := Sprite{0, 0}

	scene := Scene{}

	sequence := NewSequence(
		Step{What: MoveTo2(&s1, 1, 1), Duration: 500 * time.Millisecond},
		Step{What: MoveTo2(&s2, 10, 10), Duration: 1 * time.Second},
	)

	scene.Add(sequence.Once())

	execute(&scene, 1, func(tick int) {})

	if s1.x != 1 {
		t.Errorf("Expected the first sprite to be forwarded to 1, but got %d", s1.x)
	}

	if s2.x != 5 {
		t.Errorf("Expected the second sprite to be half way at 5, but got %d", s2.x)
	}
}

func TestCoarseSequenceShouldFastRewindIntermediateStepsOnYoYo(t *testing.T) {

	s1 := Sprite{0, 0}
	s2 := Sprite{0, 0}

	scene := Scene{}

	sequence := NewSequence(
		Step{What: MoveTo2(&s1, 1, 1), Duration: 500 * time.Millisecond},
		Step{What: MoveTo2(&s2, 10, 10), Duration: 1 * time.Second},
	)

	scene.Add(sequence.Build(YoYo))

	scene.Set(3000 * time.Millisecond) //first backward complete

	if s1.x != 0 {
		t.Errorf("Expected the first sprite to be rewinded to 0, but got %d", s1.x)
	}

	if s2.x != 0 {
		t.Errorf("Expected the second sprite to be rewinded to 0, but got %d", s2.x)
	}
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
