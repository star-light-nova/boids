package boid

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	HEIGHT float32 = 15
	WIDTH  float32 = 15

	PROTECTEDRANGE float32 = (HEIGHT + WIDTH) * 0.4
	VISUALRANGE    float32 = HEIGHT + WIDTH
	TURNFACTOR     float32 = 0.2
	MAXSPEED       float32 = 3
	MINSPEED       float32 = 2

	AVOIDFACTOR      float32 = 0.05
	MATCHFACTOR      float32 = 0.05
	CENTERFACTOR     float32 = 0.0005
	AVOIDMOUSEFACTOR float32 = 0.00025

	BORDER int32 = 50
)

var DEFAULT_COLOR *sdl.Color = &sdl.Color{R: 255, G: 255, B: 255, A: 126}
