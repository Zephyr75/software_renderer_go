package material

import (
	"image/color"
)

type MaterialType byte

const (
	FlatColor MaterialType = 0
	Texture   MaterialType = 1
)

type Material struct {
	MaterialType MaterialType
	Color        color.Color
}

func NewMaterial() Material {
	return Material{MaterialType: FlatColor, Color: color.White}
}

func ColorMaterial(color color.Color) Material {
	return Material{MaterialType: FlatColor, Color: color}
}
