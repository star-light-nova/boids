package boid

import "github.com/veandco/go-sdl2/sdl"

type Boid struct {
	texture *sdl.Texture

	X, Y int32
	W, H int32
}

func NewBoid() *Boid {
	return &Boid{
		W: 50,
		H: 50,
	}
}

func (boid *Boid) Texture() *sdl.Texture {
	return boid.texture
}

func (boid *Boid) Destroy() {
	boid.texture.Destroy()
}
