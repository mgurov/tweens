package tweens

type Movable interface {
	SetPosition(x float64, y float64)
	GetPosition() (x float64, y float64)
}

func MoveTo2(movable Movable, x float64, y float64) ChangeFunction {
	startX, startY := movable.GetPosition()
	deltaX := x - startX
	deltaY := y - startY

	return func(progress float64) {
		movable.SetPosition(
			startX+progress*deltaX,
			startY+progress*deltaY,
		)
	}
}
