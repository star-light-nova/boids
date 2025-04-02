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

func NewBoid(r *sdl.Renderer) *Boid {
	texture, err := r.CreateTexture(
		sdl.PIXELFORMAT_UNKNOWN,
		sdl.TEXTUREACCESS_STATIC,
		int32(HEIGHT),
		int32(WIDTH),
	)

	if err != nil {
		// TODO: Return err with Boid.
		panic("Couldn't create a texture for boid")
	}

	randVX := rand.Float32() + maxspeed
	randVY := rand.Float32() + minspeed

	if rand.Float32() > 0.5 {
		randVX = -randVX
	}

	if rand.Float32() > 0.5 {
		randVY = -randVY
	}

	return &Boid{
		texture: texture,

		W: WIDTH,
		H: HEIGHT,

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
