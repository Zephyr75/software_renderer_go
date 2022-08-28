package material


import (
	"image/color"
)

type MaterialType byte

const (
	FlatColor MaterialType = 0
	Texture = 1
)

type Material struct {
	materialType MaterialType
	color color.Color
}