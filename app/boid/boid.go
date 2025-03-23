package boid

import (
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
)

type Boid struct {
	texture *sdl.Texture

	X, Y float32
	W, H float32

	ProtectedRange float32
	VisualRange    float32
	AvoidFactor    float32
	MatchFactor    float32
	CenterFactor   float32

	// Velocity
	VX, VY float32
}

func NewBoid() *Boid {
	randVX := rand.Float32() + maxspeed
	randVY := rand.Float32() + minspeed

	return &Boid{
		W: 5,
		H: 5,

		VX: randVX,
		VY: randVY,

		// Ranges
		ProtectedRange: protectedRange,
		VisualRange:    visualRange,

		// Factors
		AvoidFactor:  AvoidFactor,
		MatchFactor:  MatchFactor,
		CenterFactor: CenterFactor,
	}
}

func (boid *Boid) Texture() *sdl.Texture {
	return boid.texture
}

func (boid *Boid) Destroy() {
	boid.texture.Destroy()
}
