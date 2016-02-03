package main

import (
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"

	"github.com/mgurov/tweens"
	"time"
)

const (
	Title  = "Tween demo"
	Width  = 400
	Height = 400
)

func main() {
	if err := glfw.Init(); err != nil {
		panic("Can't init glfw!" + err.Error())
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)

	window, err := glfw.CreateWindow(Width, Height, Title, nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()
	glfw.SwapInterval(1)
	gl.Init()

	initScene()

	runtime.LockOSThread()

	scene := tweens.Scene{}

	scene.AddNewSequence(
		&tweens.Step{What: onceBox.moveTo(400, 400), Duration: 10 * time.Second, Easing: tweens.EaseInOutBounce},
	)

	scene.AddNewSequence(
		&tweens.Step{What: repeatBox.moveTo(0, 200), Duration: 3 * time.Second, Easing: tweens.EaseOutQuad},
	).Repeat()

	scene.AddNewSequence(
		&tweens.Step{What: yoyoBox.moveTo(200, 0), Duration: 3 * time.Second, Easing: tweens.EaseInQuad},
	).YoYo()

	scene.AddNewSequence(
		tweens.Pause(4*time.Second),
		&tweens.Step{What: onceBox.resize(400, 400), Duration: 20 * time.Second},
	).YoYo()

	scene.AddNewSequence(
		&tweens.Step{What: wayPointsBox.moveTo(200, 200), Duration: 3 * time.Second},
		&tweens.Step{What: wayPointsBox.moveTo(200, 0), Duration: 2 * time.Second},
		&tweens.Step{What: wayPointsBox.moveTo(0, 200), Duration: 1 * time.Second},
	).Repeat()

	scene.AddNewSequence(
		&tweens.Step{What: arrow.rotate(-180), Duration: 2 * time.Second, Easing: tweens.EaseOutBounce},
		&tweens.Step{What: arrow.rotate(0), Duration: 2 * time.Second},
	).Repeat()

	arrowLegTraversalDuration := 5 * time.Second
	scene.AddNewSequence(
		&tweens.Step{What: arrow.moveTo(300, 150), Duration: arrowLegTraversalDuration},
		&tweens.Step{What: arrow.moveTo(300, 300), Duration: arrowLegTraversalDuration},
		&tweens.Step{What: arrow.moveTo(150, 300), Duration: arrowLegTraversalDuration},
		&tweens.Step{What: arrow.moveTo(150, 150), Duration: arrowLegTraversalDuration},
	).Repeat()

	// set to true for the experimental self-propelled mode where the scene gets updated in the backrgound
	// with a given frequence thus decoupling tweens from the GL drawing
	// if there are advantages going this way we'll have to check the synchronization question
	runAsync := true

	if runAsync {
		tweener := scene.RunBackground(10 * time.Millisecond)
		go func() {
			<-time.After(5 * time.Minute)
			tweener <- true
		}()
	}

	now := time.Now()

	for !window.ShouldClose() {
		if !runAsync {
			newNow := time.Now()
			delta := newNow.Sub(now)
			scene.Set(delta)
		}
		drawScene()
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

var onceBox Box = Box{X: 0.0, Y: 0.0, Width: 10, Height: 10, R: 0, G: 255, B: 255}
var repeatBox Box = Box{X: 200.0, Y: 0.0, Width: 10, Height: 10, R: 255, G: 0, B: 0}
var yoyoBox Box = Box{X: 0.0, Y: 200.0, Width: 10, Height: 10, R: 100, G: 255, B: 0}
var wayPointsBox Box = Box{X: 0.0, Y: 200.0, Width: 10, Height: 10, R: 128, G: 128, B: 0}
var arrow Arrow = Arrow{Box: Box{X: 150.0, Y: 150.0, Width: 200, Height: 50, R: 255, G: 128, B: 0}}

func initScene() {

	gl.Disable(gl.DEPTH_TEST)
	gl.Viewport(0, 0, Width, Height)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0, Width, Height, 0, 0, 1)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	// Displacement trick for exact pixelization
	gl.Translatef(0.375, 0.375, 0)
	gl.Enable(gl.TEXTURE_2D)
}

func drawScene() {
	gl.ClearColor(0.0, 0.0, 0.0, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	onceBox.Draw()
	repeatBox.Draw()
	yoyoBox.Draw()
	wayPointsBox.Draw()
	arrow.Draw()
}

type Box struct {
	// location of center
	X float64
	Y float64
	// dimension
	Width  float64
	Height float64
	// color
	R, G, B uint8
}

func (b Box) Draw() {
	w2 := b.Width / 2
	h2 := b.Height / 2
	gl.Color4f(float32(b.R)/255.0, float32(b.G)/255.0, float32(b.B)/255.0, 0)

	gl.PushMatrix()
	gl.Translated(b.X, b.Y, 0)
	gl.Begin(gl.QUADS)
	gl.Vertex2d(-w2, -h2)
	gl.Vertex2d(w2, -h2)
	gl.Vertex2d(w2, h2)
	gl.Vertex2d(-w2, h2)
	gl.End()
	gl.PopMatrix()
}

func (b *Box) Set(position []float64) {
	b.X = position[0]
	b.Y = position[1]
}

func (b *Box) Get() []float64 {
	return []float64{b.X, b.Y}
}

func (b *Box) moveTo(x, y float64) tweens.Transition {
	return tweens.LazyAccessor(b, x, y)
}

func (b *Box) resize(w, h float64) tweens.Transition {
	return tweens.LazyAccessor(&sizeAccessor{b}, w, h)
}

type sizeAccessor struct {
	subject *Box
}

func (sa *sizeAccessor) Set(size []float64) {
	sa.subject.Height = size[0]
	sa.subject.Width = size[1]
}

func (sa *sizeAccessor) Get() []float64 {
	return []float64{sa.subject.Width, sa.subject.Height}
}

type Arrow struct {
	Box

	Angle float64
}

func (a Arrow) Draw() {
	w2 := a.Width / 2
	h2 := a.Height / 2
	gl.Color4f(float32(a.R)/255.0, float32(a.G)/255.0, float32(a.B)/255.0, 0)

	gl.PushMatrix()
	gl.Translated(a.X, a.Y, 0)
	gl.Rotated(a.Angle, 0.0, 0.0, 1.0)
	gl.Begin(gl.TRIANGLES)

	gl.Vertex2d(w2, 0)
	gl.Vertex2d(-w2, h2)
	gl.Vertex2d(-w2, -h2)

	gl.End()
	gl.PopMatrix()
}

func (a *Arrow) rotate(toAngle float64) rotator {
	return rotator{toAngle, a}
}

type rotator struct {
	target float64
	angled *Arrow
}

func (r rotator) Start() tweens.TransitionCompletionFunction {
	return tweens.FunctionalAccessor(
		func() (currentState []float64) { return []float64{r.angled.Angle} },
		func(newState []float64) { r.angled.Angle = newState[0] },
		r.target,
	)
}
