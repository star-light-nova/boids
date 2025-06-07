package boid

import (
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
)

type Boid struct {
	texture *sdl.Texture

	X, Y float32
	W, H float32

	ProtectedRange   float32
	VisualRange      float32
	AvoidFactor      float32
	MatchFactor      float32
	CenterFactor     float32
	AvoidMouseFactor float32

	// Velocity
	VX, VY float32

	// Color
	Color *sdl.Color
}

func NewBoid(r *sdl.Renderer) (*Boid, error) {
	texture, err := r.CreateTexture(
		sdl.PIXELFORMAT_UNKNOWN,
		sdl.TEXTUREACCESS_STATIC,
		int32(HEIGHT),
		int32(WIDTH),
	)

	if err != nil {
		return nil, err
	}

	randVX := rand.Float32() + MAXSPEED
	randVY := rand.Float32() + MINSPEED

	if rand.Float32() > 0.5 {
		randVX = -randVX
	}

	if rand.Float32() > 0.5 {
		randVY = -randVY
	}

	color := DEFAULT_COLOR

	// Special Boid Colour Probability
	if rand.Float32() > 0.8 {
		r := uint8(rand.Uint32())
		g := uint8(rand.Uint32())
		b := uint8(rand.Uint32())

		// Putting max brightness to be sure that we will not have almost
		// invisible boids around.
		color = &sdl.Color{R: r, G: g, B: b, A: 255}
	}

	return &Boid{
		texture: texture,

		W: WIDTH,
		H: HEIGHT,

		VX: randVX,
		VY: randVY,

		// Ranges
		ProtectedRange: PROTECTEDRANGE,
		VisualRange:    VISUALRANGE,

		// Factors
		AvoidFactor:      AVOIDFACTOR,
		AvoidMouseFactor: AVOIDMOUSEFACTOR,
		MatchFactor:      MATCHFACTOR,
		CenterFactor:     CENTERFACTOR,

		Color: color,
	}, nil
}

func (boid *Boid) Texture() *sdl.Texture {
	return boid.texture
}

func (boid *Boid) Destroy() {
	boid.texture.Destroy()
}

func (boid *Boid) DefaultColor() *sdl.Color {
	return DEFAULT_COLOR
}
