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



func (l *Light) ApplyLight(v *geometry.Vector3, normal geometry.Vector3) {


	rInit, gInit, bInit, aInit := v.LightAmount.RGBA()

	rAdd := 255 - rInit
	gAdd := 255 - gInit
	bAdd := 255 - bInit
	aAdd := 255 - aInit

	var percentToApply float64

	switch l.LightType {
		case Ambient:
			percentToApply = 1
		case Directional:
			percentToApply = -normal.Dot(l.Rotation)
			if percentToApply < 0 {
				percentToApply = 0
			}
		case Point:
			percentToApply = 1
			//TODO: ambient light
	}

	rLight, gLight, bLight, aLight := l.Color.RGBA()

	rVertex := float64(rAdd) * percentToApply *  float64(rLight) / 255
	gVertex := float64(gAdd) * percentToApply *  float64(gLight) / 255
	bVertex := float64(bAdd) * percentToApply *  float64(bLight) / 255
	aVertex := float64(aAdd) * percentToApply *  float64(aLight) / 255

	if rVertex > 255 {
		rVertex = 255
	}
	if gVertex > 255 {
		gVertex = 255
	}
	if bVertex > 255 {
		bVertex = 255
	}
	if aVertex > 255 {
		aVertex = 255
	}

	v.LightAmount = color.RGBA{
		uint8(rInit + uint32(rVertex)),
		uint8(gInit + uint32(gVertex)),
		uint8(bInit + uint32(bVertex)),
		255}
}