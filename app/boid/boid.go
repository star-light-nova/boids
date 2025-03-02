package boid

import (
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
)

type Boid struct {
	texture *sdl.Texture

	X, Y int32
	W, H int32

	XDirection int32
	YDirection int32

	XStep int32
	YStep int32
}

func NewBoid() *Boid {
	var xdirection, ydirection int32 = 1, 1

	if rand.Float32() < 0.5 {
		xdirection = -xdirection
	}

	if rand.Float32() < 0.5 {
		ydirection = -ydirection
	}

	return &Boid{
		W: 50,
		H: 50,

		XStep: 1,
		YStep: 1,

		XDirection: xdirection,
		YDirection: ydirection,
	}
}

func (boid *Boid) Texture() *sdl.Texture {
	return boid.texture
}

func (boid *Boid) Destroy() {
	boid.texture.Destroy()
}
