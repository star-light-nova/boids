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

	scene, err := scene.NewScene(r)

	if err != nil {
		return fmt.Errorf("Could not initialise a scene: %v", err)
	}

	defer scene.Destroy()

	sceneErrors := scene.Run(r)

	w.Show()

	// MacOS hack to pump events.
	sdl.PumpEvents()

	runtime.LockOSThread()
	for {
		current_event := sdl.WaitEvent()

		// Global events
		switch current_event.(type) {
		case *sdl.QuitEvent:
			return nil
		case *sdl.MouseMotionEvent:
			mouse_motion_event := current_event.(*sdl.MouseMotionEvent)

			go func() { scene.MouseMotionEvents <- mouse_motion_event }()
		}

		// Scene events
		select {
		case sceneError := <-sceneErrors:
			return sceneError
		default:
		}
	}
}
