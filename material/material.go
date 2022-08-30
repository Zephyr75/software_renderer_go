package material


import (
	"image/color"
)

type MaterialType byte

const (
	FlatColor MaterialType = 0
	Texture MaterialType = 1
)

type Material struct {
	MaterialType MaterialType
	Color color.Color
}