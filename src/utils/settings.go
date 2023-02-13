package utils

import "math"

const FOV = 45
const RESOLUTION_X = 1500
const RESOLUTION_Y = 1000

func Z0() float64 {
	return (RESOLUTION_X / 2) / math.Tan((FOV / 2) * math.Pi / 180)
}

