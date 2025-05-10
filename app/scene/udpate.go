package scene

import "log"

func (scene *Scene) Update() {
	for i, boid := range scene.boids {
		var closeDx float32 = 0
		var closeDy float32 = 0
		var xVelAvg float32 = 0
		var yVelAvg float32 = 0

		var xPosAvg float32 = 0
		var yPosAvg float32 = 0

		var neighbors float32 = 0

		var sharedColor uint8 = 0
		var redAvg uint8 = 0
		var greenAvg uint8 = 0
		var blueAvg uint8 = 0

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

					redAvg += uint8(xVelAvg + xPosAvg)
					greenAvg += uint8(yVelAvg + yPosAvg)
					blueAvg += uint8((redAvg + greenAvg) / 2)

					neighbors++
					sharedColor++

					// log.Printf("R: %v, G: %v, B: %v", redAvg, greenAvg, blueAvg)
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

		if sharedColor > 0 {
			redAvg /= sharedColor
			greenAvg /= sharedColor
			blueAvg /= sharedColor

			if redAvg > 0 {
				boid.Color.R = redAvg
			}
			if greenAvg > 0 {
				boid.Color.G = greenAvg
			}
			if blueAvg > 0 {
				boid.Color.B = blueAvg
			}

			log.Println(boid.Color, redAvg, greenAvg, blueAvg, sharedColor)
		}
		// } else {
		//
		// 	boid.Color.R = 255
		// 	boid.Color.G = 255
		// 	boid.Color.B = 255
		// }

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
