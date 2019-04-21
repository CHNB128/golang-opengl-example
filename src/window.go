package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func resizeCallback(window *glfw.Window, width int, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
}

func keyCallBack(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mode glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Press {
		window.SetShouldClose(true)
	}
}

func mouseCallback(window *glfw.Window, xpos float64, ypos float64) {

}

func scrollCallback(window *glfw.Window, xoff float64, yoff float64) {

}
