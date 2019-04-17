package main

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

// ParticleEmitter ...
type ParticleEmitter struct {
	position            mgl32.Vec3
	particles           []Particle
	velocityMultiplayer float32
}

// Generate ...
func (emitter *ParticleEmitter) Generate(particleCount int) {
	for i := 0; i < particleCount; i++ {
		particle := Particle{position: emitter.position}
		emitter.particles = append(emitter.particles, particle)
	}
}

func (emitter *ParticleEmitter) update(deltaTime float32) {
	for _, particle := range emitter.particles {
		if particle.isDead() {
			particle.reset(emitter.position)
		}
		particle.update(deltaTime, emitter.velocityMultiplayer)
	}
}

func (emitter *ParticleEmitter) render(window *glfw.Window, program uint32) {
	for _, particle := range emitter.particles {
		vao := makeVao(particle.position)
		draw(vao, particle.position, window, program)
	}
}
