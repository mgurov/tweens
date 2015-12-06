package main

import (
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"

	"github.com/mgurov/tweens"
	"log"
	"time"
)

var _ log.Logger

const (
	Title = "Tween demo"
	Width = 400
	Height = 400
)

const (
	memory = iota
	tweenSize
	tweenPosition
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
	tweensManager.Add(tweens.MoveTo(&box, 400, 400, 10000))
	//positionTween.Repeat(-1, 0)

	now := time.Now()

	for !window.ShouldClose() {
		newNow := time.Now()
		delta := newNow.Sub(now).Nanoseconds() / 1000000;
		tweensManager.Set(int(delta))
		drawScene()
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

var box Box = Box{X: 0.0, Y: 0.0, Width: 10, Height: 10, R: 0, G: 255, B: 255}

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

	box.Draw()
}

type Box struct {
	// location of center
	X       float32
	Y       float32
	// dimension
	Width   float32
	Height  float32
	// color
	R, G, B uint8
}

func (b *Box) SetPosition(x int, y int) {
	b.X = float32(x)
	b.Y = float32(y)
}

func (b *Box) GetPosition() (x int, y int) {
	return int(b.X), int(y)
}

func (b Box) Draw() {
	w2 := b.Width / 2
	h2 := b.Height / 2
	gl.Color4f(float32(b.R) / 255.0, float32(b.G) / 255.0, float32(b.B) / 255.0, 0)

	gl.PushMatrix()
	gl.Translatef(b.X, b.Y, 0)
	gl.Begin(gl.QUADS)
	gl.Vertex2f(-w2, -h2)
	gl.Vertex2f(w2, -h2)
	gl.Vertex2f(w2, h2)
	gl.Vertex2f(-w2, h2)
	gl.End()
	gl.PopMatrix()
}


