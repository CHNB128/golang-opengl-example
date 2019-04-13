package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type vector3 struct {
	x float64
	y float64
	z float64
}

type vertex = vector3
type normal = vector3
type uv = vector3

type obj struct {
	vertexes []vertex
	normals  []normal
	uvs      []uv
}

func loadObj(path string) obj {
	vertexes := []vertex{}
	normals := []normal{}

	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		switch line[0] {
		case "v":
			x, err := strconv.ParseFloat(line[1], 64)
			check(err)
			y, err := strconv.ParseFloat(line[2], 64)
			check(err)
			z, err := strconv.ParseFloat(line[3], 64)
			check(err)
			vertex := vertex{x, y, z}
			vertexes = append(vertexes, vertex)
		}
	}

	return obj{vertexes: vertexes, normals: normals, uvs: nil}
}
