package scene

func (scene *Scene) Update() {
	for i, boid := range scene.boids {
		var closeDx float32 = 0
		var closeDy float32 = 0
		var xVelAvg float32 = 0
		var yVelAvg float32 = 0
		var neighbors float32 = 0

		var xPosAvg float32 = 0
		var yPosAvg float32 = 0

		for j, otherBoid := range scene.boids {
			if i == j {
				continue
			}

			dx := boid.X - otherBoid.X
			dy := boid.Y - otherBoid.Y

			if abs(dx) < boid.VisualRange && abs(dy) < boid.VisualRange {
				distance := dx*dx + dy*dy

				// Separation
				if distance < (boid.ProtectedRange * boid.ProtectedRange) {
					closeDx += boid.X - otherBoid.X
					closeDy += boid.Y - otherBoid.Y
				} else if distance < (boid.VisualRange * boid.VisualRange) {
					// Alignment
					xVelAvg += otherBoid.VX
					yVelAvg += otherBoid.VY
					// Cohesion
					xPosAvg += otherBoid.X
					yPosAvg += otherBoid.Y

					neighbors++
				}
			}
		}

		if neighbors > 0 {
			xVelAvg = xVelAvg / neighbors
			yVelAvg = yVelAvg / neighbors

			xPosAvg = xPosAvg / neighbors
			yPosAvg = yPosAvg / neighbors

			boid.VX = boid.VX + ((xPosAvg-boid.X)*boid.CenterFactor + (xVelAvg-boid.VX)*boid.MatchFactor)
			boid.VY = boid.VY + ((yPosAvg-boid.Y)*boid.CenterFactor + (yVelAvg-boid.VY)*boid.MatchFactor)
		}

		boid.VX = boid.VX + (boid.AvoidFactor * closeDx)
		boid.VY = boid.VY + (boid.AvoidFactor * closeDy)

		boid.Update()
	}
}

func abs(number float32) float32 {
	if number < 0 {
		return -number
	}

	return number
}
