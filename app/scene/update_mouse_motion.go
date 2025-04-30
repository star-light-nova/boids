package scene

import "github.com/veandco/go-sdl2/sdl"

func (scene *Scene) UpdateMouseMotion(mouse_motion_event *sdl.MouseMotionEvent) {
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

		// Calculate for mouse movement
		mouse_motion_event_X := float32(mouse_motion_event.X)
		mouse_motion_event_Y := float32(mouse_motion_event.Y)

		dx := boid.X - mouse_motion_event_X
		dy := boid.Y - mouse_motion_event_Y

		if abs(dx) < boid.VisualRange && abs(dy) < boid.VisualRange {
			distance := dx*dx + dy*dy

			// Separation
			if distance < (boid.ProtectedRange * boid.ProtectedRange) {
				closeDx += boid.X - mouse_motion_event_X
				closeDy += boid.Y - mouse_motion_event_Y
			} else if distance < (boid.VisualRange * boid.VisualRange) {
				// Alignment
				// Find velocity for mouse motion movement
				xVelAvg += mouse_motion_event_X
				yVelAvg += mouse_motion_event_Y
				// Cohesion
				xPosAvg += mouse_motion_event_X
				yPosAvg += mouse_motion_event_Y

				neighbors++
			}
		}

		if neighbors > 0 {
			xVelAvg = xVelAvg / neighbors
			yVelAvg = yVelAvg / neighbors

			xPosAvg = xPosAvg / neighbors
			yPosAvg = yPosAvg / neighbors

			boid.VX = boid.VX + ((xPosAvg-boid.X)*boid.CenterFactor + (xVelAvg-boid.VX)*boid.MatchFactor + (xPosAvg-mouse_motion_event_X)*boid.AvoidMouseFactor)
			boid.VY = boid.VY + ((yPosAvg-boid.Y)*boid.CenterFactor + (yVelAvg-boid.VY)*boid.MatchFactor + (yPosAvg-mouse_motion_event_Y)*boid.AvoidMouseFactor)
		}

		boid.VX = boid.VX + (boid.AvoidFactor * closeDx)
		boid.VY = boid.VY + (boid.AvoidFactor * closeDy)

		boid.Update()
	}
}
