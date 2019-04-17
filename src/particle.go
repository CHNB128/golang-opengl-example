package main

import (
	"github.com/go-gl/mathgl/mgl32"
)

// Particle ...
type Particle struct {
	position     mgl32.Vec3
	direction    mgl32.Vec3
	ttl          float32
	velocity     mgl32.Vec3
	acceleration mgl32.Vec3
	color        string
}

func (particle *Particle) update(deltaTime float32, multiplayer float32) {
	particle.velocity.Mul(multiplayer * deltaTime)
	particle.ttl = particle.ttl - deltaTime
}

func (particle *Particle) isDead() bool {
	return particle.ttl <= 0.0
}

func (particle *Particle) reset(position mgl32.Vec3) {
	particle.position = position
	particle.direction = mgl32.Vec3{0, 0, 0}
	particle.ttl = 0.0
	particle.velocity = mgl32.Vec3{0, 0, 0}
	particle.acceleration = mgl32.Vec3{0, 0, 0}
	particle.color = ""
}
