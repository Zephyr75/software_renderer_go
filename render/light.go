package render


import (
	"mesh/vector3"
)

type LightType byte

const (
	FlatColor MaterialType = 0
	RTexture = 1


type Light struct {
	position vector3.Vector3
	direction vector3.Vector3
	lightType LightType
	color color.Color
	length float64
}