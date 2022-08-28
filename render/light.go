package render


import (
	"image/color"
	"overdrive/mesh"
)

type LightType byte

const (
	Directional LightType = 0
	Point = 1
	Ambient = 2
)

type Light struct {
	position mesh.Vector3
	direction mesh.Vector3
	lightType LightType
	color color.Color
	length float64
}