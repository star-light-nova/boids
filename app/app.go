package app

import (
	"boids/app/config"
	"boids/app/scene"
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func Run() error {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return fmt.Errorf("Could not init App: %v", err)
	}

	// Window
	w, r, err := sdl.CreateWindowAndRenderer(
		config.DEFAULT_WINDOW_WIDTH,
		config.DEFAULT_WINDOW_HEIGHT,
		sdl.WINDOW_SHOWN,
	)
	defer w.Destroy()
	defer r.Destroy()

	if err != nil {
		return fmt.Errorf("Could not create either a window or renderer or both: %v", err)
	}

	scene := scene.NewScene()
	defer scene.Destroy()

	if err := scene.Paint(r); err != nil {
		return fmt.Errorf("Scene can't paint: %v", err)
	}

	w.Show()

	// MacOs hack to pump events.
	sdl.PumpEvents()

	for {
		event := sdl.WaitEvent()

		switch event.(type) {
		case *sdl.QuitEvent:
			return nil
		default:
		}
	}
}
