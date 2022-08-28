package utilities

import "math"

const FOV = 45
const RESOLUTION_X = 1000
const RESOLUTION_Y = 600

func z0() float64 {
	return (RESOLUTION_X / 2) / math.Tan((FOV / 2) * math.Pi / 180)
}