package boid

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func (boid *Boid) Paint(r *sdl.Renderer) error {
	rect := &sdl.FRect{
		X: boid.X,
		Y: boid.Y,
		W: boid.W,
		H: boid.H,
	}

	texture, err := r.CreateTexture(
		sdl.PIXELFORMAT_UNKNOWN,
		sdl.TEXTUREACCESS_STATIC,
		int32(boid.W),
		int32(boid.H),
	)

	if err != nil {
		return fmt.Errorf("Couldn't create texture for boid: %v", err)
	}

	boid.texture = texture

	err = r.SetDrawColor(255, 255, 255, 126)

	if err != nil {
		return fmt.Errorf("Couldn't set draw color for boid: %v", err)
	}

	if err := r.CopyF(boid.Texture(), nil, rect); err != nil {
		return fmt.Errorf("Couldn't copy boid: %v", err)
	}

	if err := r.FillRectF(rect); err != nil {
		return fmt.Errorf("Couldn't fill react with color: %v", err)
	}

	return nil
}
