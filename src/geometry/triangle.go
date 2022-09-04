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

	//TODO : fix weights not corresponding to the right vertex

	vertices := []Vector3{t.A, t.B, t.C}

	points := make([]Point, 3)

	for i, v := range vertices {
		points[i] = v.Converted()
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			if points[j].Y > points[j+1].Y {
				points[j], points[j+1] = points[j+1], points[j]
				vertices[j], vertices[j+1] = vertices[j+1], vertices[j]
			}
		}
	}

	var p0, p1, p2 = points[0], points[1], points[2]
	var v0, v1, v2 = vertices[0], vertices[1], vertices[2]
	var top, mid, bottom = p0.Y, p1.Y, p2.Y

	wg := sync.WaitGroup{}

	rT, gT, bT, _ := t.Material.Color.RGBA()
	
	var width, height float32

	if t.Material.MaterialType == material.Texture {
		width = float32(t.Material.Image.Bounds().Max.X)
		height = float32(t.Material.Image.Bounds().Max.Y)
	}

	r0i, g0i, b0i, _ := v0.LightAmount.RGBA()
	r1i, g1i, b1i, _ := v1.LightAmount.RGBA()
	r2i, g2i, b2i, _ := v2.LightAmount.RGBA()
	r0 := float32(r0i) / 16842495
	g0 := float32(g0i) / 16842495
	b0 := float32(b0i) / 16842495
	r1 := float32(r1i) / 16842495
	g1 := float32(g1i) / 16842495
	b1 := float32(b1i) / 16842495
	r2 := float32(r2i) / 16842495
	g2 := float32(g2i) / 16842495
	b2 := float32(b2i) / 16842495

	dist0 := v0.Distance(ZeroVector())
	dist1 := v1.Distance(ZeroVector())
	dist2 := v2.Distance(ZeroVector())

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

				weight0 := float32(num1) / float32(denom1)
				weight1 := float32(num2) / float32(denom2)
				weight2 := 1 - weight0 - weight1
				
				// fmt.Println(num1, denom1, num2, denom2)
				// fmt.Println(weight0, weight1, weight2)
				// fmt.Println("-----------------")

				r := weight0*r0 + weight1*r1 + weight2*r2
				g := weight0*g0 + weight1*g1 + weight2*g2
				b := weight0*b0 + weight1*b1 + weight2*b2

				z := weight0*dist0 + weight1*dist1 + weight2*dist2
				if x >= 0 && x < utilities.RESOLUTION_X && y >= 0 && y < utilities.RESOLUTION_Y {
					if z < zBuffer[y*utilities.RESOLUTION_X+x] || zBuffer[y*utilities.RESOLUTION_X+x] < 0 {
						zBuffer[y*utilities.RESOLUTION_X+x] = z
						
						if t.Material.MaterialType == material.Texture {
							u := weight0*v0.U + weight1*v1.U + weight2*v2.U
							v := weight0*v0.V + weight1*v1.V + weight2*v2.V
							u *= width
							v *= height
							rT, gT, bT, _ = t.Material.Image.At(int(u), int(v)).RGBA()
						} 
						img.Set(x, y, color.RGBA{
							uint8(r * float32(rT)), 
							uint8(g * float32(gT)), 
							uint8(b * float32(bT)), 
							255})
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
