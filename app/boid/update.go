package boid

import (
	"boids/app/config"
	"math"
)

func (boid *Boid) Update() {
	// Top
	if boid.Y < float32(border) {
		boid.VY = boid.VY + turnFactor
	}

	// Right
	if boid.X+boid.W > float32(config.DEFAULT_WINDOW_WIDTH-border) {
		boid.VX = boid.VX - turnFactor
	}

	// Left
	if boid.X < float32(border) {
		boid.VX = boid.VX + turnFactor
	}

	// Bottom
	if boid.Y+boid.H > float32(config.DEFAULT_WINDOW_HEIGHT-border) {
		boid.VY = boid.VY - turnFactor
	}

	// Speed limits
	speed := float32(math.Sqrt(float64(boid.VX*boid.VX + boid.VY*boid.VY)))

	if speed < minspeed {
		boid.VX = (boid.VX / speed) * minspeed
		boid.VY = (boid.VY / speed) * minspeed
	}

	if speed > maxspeed {
		boid.VX = (boid.VX / speed) * maxspeed
		boid.VY = (boid.VY / speed) * maxspeed
	}

	boid.X = boid.X + boid.VX
	boid.Y = boid.Y + boid.VY
}
