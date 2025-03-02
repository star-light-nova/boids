package boid

func (boid *Boid) Update() {
	if boid.X+boid.W >= 800 || boid.X <= 0 {
		boid.XDirection = -boid.XDirection
	}

	if boid.Y+boid.H >= 600 || boid.Y <= 0 {
		boid.YDirection = -boid.YDirection
	}

	boid.X += boid.XDirection
	boid.Y += boid.YDirection
}
