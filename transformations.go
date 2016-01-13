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

type RotateAble interface {
	GetAngle() (rad float64)
	SetAngle(rad float64) ()
}

func Rotate(movable RotateAble, targetAngle float64) ChangeFunction {
	start := movable.GetAngle()
	delta := targetAngle - start

	return func(progress float64) {
		movable.SetAngle(start + progress * delta)
	}
}

