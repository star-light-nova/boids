package boid

import "github.com/veandco/go-sdl2/sdl"

const (
	protectedRange float32 = 5
	visualRange    float32 = 20
	turnFactor     float32 = 0.2
	maxspeed       float32 = 3
	minspeed       float32 = 2

	AvoidFactor      float32 = 0.05
	MatchFactor      float32 = 0.05
	CenterFactor     float32 = 0.0005
	AvoidMouseFactor float32 = 0.00025

	border int32 = 50

	HEIGHT float32 = 5
	WIDTH  float32 = 5
)

var DEFAULT_COLOR *sdl.Color = &sdl.Color{R: 255, G: 255, B: 255, A: 126}
