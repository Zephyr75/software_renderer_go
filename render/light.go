package render


import (
	"image/color"
	"overdrive/geometry"
	"fmt"
)

type LightType byte

const (
	Directional LightType = 0
	Point LightType = 1
	Ambient LightType = 2
)

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


	fmt.Println("---------------------")
	switch l.LightType {
		case Ambient:
			percentToApply = 1
		case Directional:
			lightRotation := l.Rotation
			lightRotation.Normalize()
			percentToApply = normal.Dot(lightRotation)
			
			fmt.Println(percentToApply)
			if percentToApply < 0 {
				percentToApply *= -1
			}
		case Point:
			percentToApply = 1
			//TODO: ambient light
	}

	rLight, gLight, bLight, _ := l.Color.RGBA()
	rLight /= 257
	gLight /= 257
	bLight /= 257

	fmt.Println(rInit)
	fmt.Println(rAdd)
	fmt.Println(rLight)
	fmt.Println(percentToApply)


	rVertex := rInit + uint32(float64(rAdd) * percentToApply *  float64(rLight) / 255)
	gVertex := gInit + uint32(float64(gAdd) * percentToApply *  float64(gLight) / 255)
	bVertex := bInit + uint32(float64(bAdd) * percentToApply *  float64(bLight) / 255)

	fmt.Println("rVertex: ", rVertex)
	fmt.Println("gVertex: ", gVertex)
	fmt.Println("bVertex: ", bVertex)

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