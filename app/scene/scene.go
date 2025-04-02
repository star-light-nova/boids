package scene

import (
	"boids/app/boid"
	"math/rand"

	ap "boids/app/config"
	"github.com/veandco/go-sdl2/sdl"
)

type Scene struct {
	boids []*boid.Boid

	X, Y int32
	W, H int32
}

func NewScene(r *sdl.Renderer) *Scene {
	boids := []*boid.Boid{}

	for range 2000 {
		boid := boid.NewBoid(r)

		randX := float32(rand.Int31n(ap.DEFAULT_WINDOW_WIDTH - int32(boid.W)))
		randY := float32(rand.Int31n(ap.DEFAULT_WINDOW_HEIGHT - int32(boid.H)))

		boid.X, boid.Y = randX, randY

		boids = append(boids, boid)
	}

	return &Scene{
		boids: boids,
	}
}

func (scene *Scene) Destroy() {
	for _, boid := range scene.boids {
		boid.Destroy()
	}
}
