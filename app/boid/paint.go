package boid

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

var SimpleRect = &sdl.FRect{
	X: 0,
	Y: 0,
	W: 0,
	H: 0,
}

func (boid *Boid) Paint(r *sdl.Renderer) error {
	SimpleRect.X = boid.X
	SimpleRect.Y = boid.Y
	SimpleRect.W = boid.W
	SimpleRect.H = boid.H

	rect := SimpleRect

	err := r.SetDrawColor(boid.Color.R, boid.Color.G, boid.Color.B, boid.Color.A)

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
