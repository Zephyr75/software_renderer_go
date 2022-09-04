package geometry

import (
	"image"
	"image/color"
	"overdrive/src/material"
	"overdrive/src/utilities"

	"sync"
)

type Triangle struct {
	A        Vector3
	B        Vector3
	C        Vector3
	Material material.Material
}

func NewTriangle(a, b, c Vector3) Triangle {
	return Triangle{a, b, c, material.NewMaterial()}
}

func (t Triangle) Normal() Vector3 {
	v1 := t.B.Sub(t.A)
	v1.Normalize()
	v2 := t.C.Sub(t.A)
	v2.Normalize()
	return v1.Cross(v2)
}

func (t Triangle) Average() Vector3 {
	return t.A.Add(t.B).Add(t.C).Div(3)
}

func (t *Triangle) Draw(img *image.RGBA, zBuffer []float32) {

	// fmt.Println(t.Material.Color.RGBA())

	// fmt.Println("2")
	// fmt.Println(t.A.LightAmount.RGBA())

	vertices := []Vector3{t.A, t.B, t.C}

	points := make([]Point, 3)

	for i, v := range vertices {
		points[i] = v.Converted()
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			if points[j].Y > points[j+1].Y {
				points[j], points[j+1] = points[j+1], points[j]
			}
		}
	}

	var p0, p1, p2 = points[0], points[1], points[2]
	var top, mid, bottom = p0.Y, p1.Y, p2.Y

	wg := sync.WaitGroup{}

	rT, gT, bT, _ := t.Material.Color.RGBA()

	rAi, gAi, bAi, _ := t.A.LightAmount.RGBA()
	rBi, gBi, bBi, _ := t.B.LightAmount.RGBA()
	rCi, gCi, bCi, _ := t.C.LightAmount.RGBA()
	rA := float32(rAi*rT) / 16842495
	gA := float32(gAi*gT) / 16842495
	bA := float32(bAi*bT) / 16842495
	rB := float32(rBi*rT) / 16842495
	gB := float32(gBi*gT) / 16842495
	bB := float32(bBi*bT) / 16842495
	rC := float32(rCi*rT) / 16842495
	gC := float32(gCi*gT) / 16842495
	bC := float32(bCi*bT) / 16842495

	distA := t.A.Distance(ZeroVector())
	distB := t.B.Distance(ZeroVector())
	distC := t.C.Distance(ZeroVector())

	for y := top; y < bottom; y++ {

		wg.Add(1)

		go func(y int) {
			var min int
			var max int
			if y < mid {
				a := f(p0, p1, y)
				b := f(p0, p2, y)
				min = utilities.Min(a, b)
				max = utilities.Max(a, b)
			} else {
				a := f(p1, p2, y)
				b := f(p0, p2, y)
				min = utilities.Min(a, b)
				max = utilities.Max(a, b)
			}

			for x := min; x < max; x++ {

				num1 := (p1.Y-p2.Y)*(x-p2.X) + (p2.X-p1.X)*(y-p2.Y)
				num2 := (p2.Y-p0.Y)*(x-p2.X) + (p0.X-p2.X)*(y-p2.Y)

				denom1 := (p1.Y-p2.Y)*(p0.X-p2.X) + (p2.X-p1.X)*(p0.Y-p2.Y)
				if denom1 == 0 {
					denom1 = 1
				}
				denom2 := (p1.Y-p2.Y)*(p0.X-p2.X) + (p2.X-p1.X)*(p0.Y-p2.Y)
				if denom2 == 0 {
					denom2 = 1
				}

				weight0 := float32(num1 / denom1)
				weight1 := float32(num2 / denom2)
				weight2 := 1 - weight0 - weight1

				r := weight0*rA + weight1*rB + weight2*rC
				g := weight0*gA + weight1*gB + weight2*gC
				b := weight0*bA + weight1*bB + weight2*bC

				z := weight0*distA + weight1*distB + weight2*distC
				if x >= 0 && x < utilities.RESOLUTION_X && y >= 0 && y < utilities.RESOLUTION_Y {
					if z < zBuffer[y*utilities.RESOLUTION_X+x] || zBuffer[y*utilities.RESOLUTION_X+x] < 0 {
						zBuffer[y*utilities.RESOLUTION_X+x] = z
						img.Set(x, y, color.RGBA{uint8(r), uint8(g), uint8(b), 255})
					}
				}
			}
			wg.Done()
		}(y)
	}

	wg.Wait()

}

func f(start, end Point, y int) int {
	height := end.Y - start.Y
	if height == 0 {
		height = 1
	}
	return start.X + (end.X-start.X)*(y-start.Y)/height
}
