package boid

import (
	"boids/app/config"
	"math"
)

func (boid *Boid) Update() {
	// Top
	if boid.Y < float32(BORDER) {
		boid.VY = boid.VY + TURNFACTOR
	}

	// Right
	if boid.X+boid.W > float32(config.DEFAULT_WINDOW_WIDTH-BORDER) {
		boid.VX = boid.VX - TURNFACTOR
	}

	// Left
	if boid.X < float32(BORDER) {
		boid.VX = boid.VX + TURNFACTOR
	}

	// Bottom
	if boid.Y+boid.H > float32(config.DEFAULT_WINDOW_HEIGHT-BORDER) {
		boid.VY = boid.VY - TURNFACTOR
	}

	// Speed limits
	speed := float32(math.Sqrt(float64(boid.VX*boid.VX + boid.VY*boid.VY)))

	if speed < MINSPEED {
		boid.VX = (boid.VX / speed) * MINSPEED
		boid.VY = (boid.VY / speed) * MINSPEED
	}

	if speed > MAXSPEED {
		boid.VX = (boid.VX / speed) * MAXSPEED
		boid.VY = (boid.VY / speed) * MAXSPEED
	}

	boid.X = boid.X + boid.VX
	boid.Y = boid.Y + boid.VY
}
