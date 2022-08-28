package render


import (
	"image/color"
	"overdrive/geometry"
)

type LightType byte

const (
	Directional LightType = 0
	Point = 1
	Ambient = 2
)

type Light struct {
	Position geometry.Vector3
	Rotation geometry.Vector3
	LightType LightType
	Color color.Color
	Length float64
}



func (l *Light) ApplyLight(v geometry.Vector3, normal geometry.Vector3) {
	v.LightAmount = color.White
	//TODO
}