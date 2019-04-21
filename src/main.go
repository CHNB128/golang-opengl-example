package main

import (
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
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

// OpenGL
func init() {
	if err := gl.Init(); err != nil {
		panic(err)
	}
}

// GLFW
func init() {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
}

// GLFW Window
func init() {
	_window, err := glfw.CreateWindow(500, 500, "Test", nil, nil)
	if err != nil {
		panic(err)
	}
	window = _window
	window.MakeContextCurrent()
	window.SetKeyCallback(keyCallBack)
	window.SetFramebufferSizeCallback(resizeCallback)
	window.SetCursorPosCallback(mouseCallback)
	window.SetScrollCallback(scrollCallback)
}

// OpenGL Program
func init() {
	program := gl.CreateProgram()
	AttachShaders(program, "./src/shaders")
	gl.LinkProgram(program)
}

func main() {
	defer glfw.Terminate()

	triangle := []float32{
		0, 0.5, 0, // top
		-0.5, -0.5, 0, // left
		0.5, -0.5, 0, // right
	}
	triangleVAO := makeVao(triangle)
	triangleRender := Renderable{vao: triangleVAO, vertexes: triangle}
	renderQueue = append(renderQueue, triangleRender)

	loop()
}

func loop() {
	for !window.ShouldClose() {
		render(renderQueue, window, program)
	}
}
