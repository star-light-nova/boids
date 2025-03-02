package scene

func (scene *Scene) Update() {
	for _, boid := range scene.boids {
		boid.Update()
	}
}
