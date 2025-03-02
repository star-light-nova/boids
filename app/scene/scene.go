package scene

import (
	"boids/app/boid"
	"math/rand"
)

type Scene struct {
	boids []*boid.Boid

	X, Y int32
	W, H int32
}

func NewScene() *Scene {
	boids := []*boid.Boid{}

	for range 10 {
		boid := boid.NewBoid()

		randX := rand.Int31n(800 - boid.W)
		randY := rand.Int31n(600 - boid.H)

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
