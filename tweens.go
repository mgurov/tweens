package tweens

import (
	"math"
	"time"
	_ "fmt"
	_ "log"
)

type Scene struct {
	items []Setter
}

func (s *Scene) AddTransition(what ChangeFunction, how How) {
	s.Add(&transition{setterFunctions: []ChangeFunction{what}, tickNormalizer: how.tickNormalizationFun()})
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

type ChangeFunction func(complete float64)
type TickNormalizationFunction func(tick time.Duration) float64

type How struct {
	Duration   time.Duration
	Easing     Easing
	Repetition Lifespan
}

func (t How) tickNormalizationFun() TickNormalizationFunction {

	if nil == t.Easing {
		t.Easing = Linear
	}

	if nil == t.Repetition {
		t.Repetition = Once
	}

	return func(tick time.Duration) float64 {
		var completed float64
		if tick <= 0 {
			//shift the tiniest bit to make downstream functions happier. Kind of a special case, although might need rethinking.
			completed = math.Nextafter(0, 1)
		} else {
			completed = float64(tick) / float64(t.Duration)
		}

		return t.Easing(t.Repetition(completed))
	}
}

type transition struct {
	setterFunctions []ChangeFunction //TODO: I wonder whether how would it work with pointers and whether it would make any sense
	tickNormalizer  TickNormalizationFunction
}

func (fs *transition) Set(tick time.Duration) {
	normalized := fs.tickNormalizer(tick)
	for _, fun := range fs.setterFunctions {
		fun(normalized)
	}
}
