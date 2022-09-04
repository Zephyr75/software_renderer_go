package render


import (
	"image/color"
	"overdrive/geometry"
)

//Enum for light types
type LightType byte
const (
	Directional LightType = 0
	Point LightType = 1
	Ambient LightType = 2
)

//Light component in the scene of given type and color
type Light struct {
	Position geometry.Vector3
	Rotation geometry.Vector3
	LightType LightType
	Color color.Color
	Length float64
}



func (l Light) ApplyLight(v *geometry.Vector3, normal geometry.Vector3) {

	rInit, gInit, bInit, _ := v.LightAmount.RGBA()

	rInit /= 257
	gInit /= 257
	bInit /= 257

	rAdd := 255 - rInit
	gAdd := 255 - gInit
	bAdd := 255 - bInit

	var percentToApply float64


	switch l.LightType {
		case Ambient:
			percentToApply = 1
		case Directional:
			l.Rotation.Normalize()
			percentToApply = normal.Dot(l.Rotation)
		case Point:
			direction := l.Position.Sub(*v)
			dim := 1 - direction.Norm() / l.Length
			direction.Normalize()
			percentToApply = normal.Dot(direction) * dim
	}
	if percentToApply < 0 {
		percentToApply = 0
	}

	rLight, gLight, bLight, _ := l.Color.RGBA()
	rLight /= 257
	gLight /= 257
	bLight /= 257

	rVertex := rInit + uint32(float64(rAdd) * percentToApply *  float64(rLight) / 255)
	gVertex := gInit + uint32(float64(gAdd) * percentToApply *  float64(gLight) / 255)
	bVertex := bInit + uint32(float64(bAdd) * percentToApply *  float64(bLight) / 255)

	if rVertex > 255 {
		rVertex = 255
	}
	if gVertex > 255 {
		gVertex = 255
	}
	if bVertex > 255 {
		bVertex = 255
	}

	v.LightAmount = color.RGBA{
		uint8(rVertex),
		uint8(gVertex),
		uint8(bVertex),
		255}
}