package tweens

import (
	"time"
)

type Scene struct {
	changes []Change
}

// Changes progress over the time from the moment zero which is normally when the Scene has started
type Change interface {
	Progress(tick time.Duration)
}

func (s *Scene) Add(newSetter Change) {
	s.changes = append(s.changes, newSetter)
}

func (s *Scene) AddNewSequence(steps ...*Step) *Sequence {
	sequence := NewSequence(steps...)
	s.Add(sequence)
	return sequence
}

// Sets the timestamp
func (s *Scene) Set(t time.Duration) {
	for _, m := range s.changes {
		m.Progress(t)
	}
}

func (s *Scene) RunInfinitely(tickPrecision time.Duration) {
	s.RunUntilStopped(tickPrecision, nil)
}

func (s *Scene) RunBackground(tickPrecision time.Duration) (quit chan bool) {
	quit = make(chan bool)
	go s.RunUntilStopped(tickPrecision, quit)
	return quit
}

func (s *Scene) RunUntilStopped(tickPrecision time.Duration, quit chan bool) {
	startTime := time.Now()
	ticker := time.NewTicker(tickPrecision)
	for {
		select {
		case now := <-ticker.C:
			//TODO: report as a bug to https://github.com/go-lang-plugin-org/go-lang-idea-plugin/issues?q=is%3Aissue+label%3A%22type+inference%22+is%3Aopen
			// I mean now := range time.Tick(precision)
			s.Set(now.Sub(startTime))
		case <-quit:
			ticker.Stop()
			return
		}
	}
}
