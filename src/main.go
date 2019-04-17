package main

import (
	"runtime"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

var (
	window      *glfw.Window
	program     uint32
	renderQueue RenderQueue
	emitter     ParticleEmitter
)

func init() {
	runtime.LockOSThread()
}

func main() {
	lwindow, err := CreateWindow(400, 500, "test")
	check(err)
	window = lwindow

	defer glfw.Terminate()
	program = initOpenGL()

	// model := NewModel("./assets/dolphin.obj")
	// renderableVertices := model.GetRenderableVertices()
	// vao := makeVao(renderableVertices)
	emitter = ParticleEmitter{position: mgl32.Vec3{0, 0, 0}, velocityMultiplayer: 0.0001, particles: []Particle{}}
	emitter.Generate(100)
	loop()
}

func loop() {
	for !window.ShouldClose() {
		glfw.PollEvents()
		emitter.render(window, program)
		render(renderQueue, window, program)
		window.SwapBuffers()
	}
}
