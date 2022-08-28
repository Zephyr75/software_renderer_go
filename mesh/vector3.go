package mesh

import (
	"image/color"
	"render/material"
)

type Vector3 struct {
	x float64
	y float64
	z float64
	lightAmount color.Color
	material Material
}