package app

import (
	"boids/app/config"
	"boids/app/scene"
	"fmt"
	"runtime"

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

	scene := scene.NewScene(r)
	defer scene.Destroy()

	// events := make(chan sdl.Event)
	// defer close(events)

	sceneErrors := scene.Run(r)

	w.Show()

	// MacOS hack to pump events.
	sdl.PumpEvents()

	runtime.LockOSThread()
	for {
		// Global events
		switch sdl.WaitEvent().(type) {
		case *sdl.QuitEvent:
			return nil
		}

		// Scene events
		select {
		case sceneError := <-sceneErrors:
			return sceneError
		default:
		}
	}
}
