package main

import (
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"

	"github.com/mgurov/tweens"
	"time"
)

const (
	Title = "Tween demo"
	Width = 400
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

	/*
		sizeTween := tween.ToAccessor(&box, &ResizeBox{tweenSize}, 5).Target2(100, 50)
		//TODO: why didn't chaining work?
		sizeTween.RepeatYoyo(-1, 0)
		sizeTween.Start()
	*/

	tweensManager := tweens.Scene{}
	tweensManager.Add(tweens.MoveTo(&onceBox, 400, 400, time.Duration(10) * time.Second, tweens.EaseInOutBounce))
	tweensManager.Add(tweens.MoveToRepeat(&repeatBox, 0, 200, time.Duration(3) * time.Second, tweens.EaseOutQuad, tweens.Repeat))
	tweensManager.Add(tweens.MoveToRepeat(&yoyoBox, 200, 0, time.Duration(3) * time.Second, tweens.EaseInQuad, tweens.YoYo))

	now := time.Now()

	for !window.ShouldClose() {
		newNow := time.Now()
		delta := newNow.Sub(now);
		tweensManager.Set(delta)
		drawScene()
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

var onceBox Box = Box{X: 0.0, Y: 0.0, Width: 10, Height: 10, R: 0, G: 255, B: 255}
var repeatBox Box = Box{X: 200.0, Y: 0.0, Width: 10, Height: 10, R: 255, G: 0, B: 0}
var yoyoBox Box = Box{X: 0.0, Y: 200.0, Width: 10, Height: 10, R: 100, G: 255, B: 0}

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
}

type Box struct {
	// location of center
	X       float64
	Y       float64
	// dimension
	Width   float64
	Height  float64
	// color
	R, G, B uint8
}

func (b *Box) SetPosition(x float64, y float64) {
	b.X = x
	b.Y = y
}

func (b *Box) GetPosition() (x float64, y float64) {
	return b.X, b.Y
}

func (b Box) Draw() {
	w2 := b.Width / 2
	h2 := b.Height / 2
	gl.Color4f(float32(b.R) / 255.0, float32(b.G) / 255.0, float32(b.B) / 255.0, 0)

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