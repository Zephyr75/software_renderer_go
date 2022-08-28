package render


import (
	"image/color"
)

type MaterialType byte

const (
	FlatColor MaterialType = 0
	RTexture = 1
)

type Material struct {
	materialType MaterialType
	color color.Color
}