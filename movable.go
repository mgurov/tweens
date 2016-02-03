package tweens

type Movable interface {
	SetPosition(x float64, y float64)
	GetPosition() (x float64, y float64)
}

type movableAccessor struct {
	subject Movable
}

func (ma movableAccessor) Get() []float64 {
	x, y := ma.subject.GetPosition()
	return []float64{x, y}
}

func (ma movableAccessor) Set(newState []float64) {
	ma.subject.SetPosition(newState[0], newState[1])
}

func MoveTo2(movable Movable, x float64, y float64) Transition {
	return LazyAccessor(movableAccessor{movable}, x, y)
}
