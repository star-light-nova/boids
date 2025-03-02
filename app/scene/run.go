package scene

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

func (scene *Scene) Run(r *sdl.Renderer) <-chan error {
	errc := make(chan error)

	go func() {
		defer close(errc)

		tick := time.Tick(16 * time.Millisecond)

		for {
			select {
			case <-tick:
				scene.Update()

				err := scene.Paint(r)
				if err != nil {
					errc <- fmt.Errorf("Couldn't paint scene in Run: %v", err)
				}
			}
		}

	}()

	return nil
}
