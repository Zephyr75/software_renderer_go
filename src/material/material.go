package material

import (
	"image"
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
	Image        image.Image
}

func WhiteMaterial() Material {
	return Material{MaterialType: FlatColor, Color: color.White, Image: nil}
}

func ColorMaterial(color color.Color) Material {
	return Material{MaterialType: FlatColor, Color: color, Image: nil}
}

func TextureMaterial(image image.Image) Material {
	return Material{MaterialType: Texture, Color: color.White, Image: image}
}
