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
	position geometry.Vector3
	direction geometry.Vector3
	lightType LightType
	color color.Color
	length float64
}



func (l *Light) ApplyLight(v geometry.Vector3, normal geometry.Vector3) {
	v.LightAmount = color.White
	//TODO
}