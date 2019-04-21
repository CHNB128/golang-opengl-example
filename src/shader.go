package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

// AttachShaders attach shader from src/shaders folder
func AttachShaders(program uint32, shaderFolderDilePath string) {
	files, err := ioutil.ReadDir(shaderFolderDilePath)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		bytes, err := ioutil.ReadFile(path.Join(shaderFolderDilePath, file.Name()))
		if err != nil {
			panic(err)
		}
		source := string(bytes)
		attachShader(program, source)
	}
}

func attachShader(programm uint32, shaderSource string) {
	shader, err := compileShader(shaderSource+"\x00", gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}
	gl.AttachShader(programm, shader)
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
