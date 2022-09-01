package geometry

import (
	"image"
	"image/color"
	"overdrive/material"
	"overdrive/utilities"

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

	
	rA, gA, bA, _ := t.A.LightAmount.RGBA()
	rB, gB, bB, _ := t.B.LightAmount.RGBA()
	rC, gC, bC, _ := t.C.LightAmount.RGBA()
	rAf := float32(rA) / 16842495
	gAf := float32(gA) / 16842495
	bAf := float32(bA) / 16842495
	rBf := float32(rB) / 16842495
	gBf := float32(gB) / 16842495
	bBf := float32(bB) / 16842495
	rCf := float32(rC) / 16842495
	gCf := float32(gC) / 16842495
	bCf := float32(bC) / 16842495
	
	rT, gT, bT, _ := t.Material.Color.RGBA()
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
				if a < b {
					min = a
					max = b
				} else {
					min = b
					max = a
				}
			} else {
				a := f(p1, p2, y)
				b := f(p0, p2, y)
				if a < b {
					min = a
					max = b
				} else {
					min = b
					max = a
				}
			}

			for x := min; x < max; x++ {

				num1 := (p1.Y-p2.Y) * (x-p2.X) + (p2.X-p1.X) * (y-p2.Y)
				num2 := (p2.Y-p0.Y) * (x-p2.X) + (p0.X-p2.X) * (y-p2.Y)

				denum1 := (p1.Y-p2.Y) * (p0.X-p2.X) + (p2.X-p1.X) * (p0.Y-p2.Y)
				if denum1 == 0 {
					denum1 = 1
				}
				denum2 := (p1.Y-p2.Y) * (p0.X-p2.X) + (p2.X-p1.X) * (p0.Y-p2.Y)
				if denum2 == 0 {
					denum2 = 1
				}

				weight0 := float32(num1 / denum1)
				weight1 := float32(num2 / denum2)
				weight2 := 1 - weight0 - weight1


				r := weight0 * rAf + weight1 * rBf + weight2 * rCf
				g := weight0 * gAf + weight1 * gBf + weight2 * gCf
				b := weight0 * bAf + weight1 * bBf + weight2 * bCf


				z := weight0 * distA + weight1 * distB + weight2 * distC
				if x >= 0 && x < utilities.RESOLUTION_X && y >= 0 && y < utilities.RESOLUTION_Y {
					if z < zBuffer[y*utilities.RESOLUTION_X+x] || zBuffer[y*utilities.RESOLUTION_X+x] < 0 {
						zBuffer[y*utilities.RESOLUTION_X+x] = z
						img.Set(x, y, color.RGBA{
							uint8(r * float32(rT)),
							uint8(g * float32(gT)),
							uint8(b * float32(bT)), 255})
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
