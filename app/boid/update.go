package boid

import "math/rand"

func (boid *Boid) Update() {
	randX := rand.Int31n(800)
	randY := rand.Int31n(600)

	boid.X, boid.Y = randX, randY
}
