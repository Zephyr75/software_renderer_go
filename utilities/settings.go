package utilities

import "math"

const FOV = 45
const RESOLUTION_X = 1280
const RESOLUTION_Y = 720

func Z0() float64 {
	return (RESOLUTION_X / 2) / math.Tan((FOV / 2) * math.Pi / 180)
}