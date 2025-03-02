package scene

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func (scene *Scene) Paint(r *sdl.Renderer) error {
	if err := r.SetDrawColor(0, 0, 0, 0); err != nil {
		return fmt.Errorf("Couldn't setup scene color: %v", err)
	}

	r.Clear()

	for i, boid := range scene.boids {
		err := boid.Paint(r)

		if err != nil {
			return fmt.Errorf("Couldn't paint boid: #%d: %v", i, boid)
		}
	}

	r.Present()

	return nil
}
