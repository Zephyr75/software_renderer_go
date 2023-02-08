package draw

import (
	"overdrive/src/geometry"
	"overdrive/src/material"
	"overdrive/src/render"
	"overdrive/src/utils"
	"sync"
)

func Draw(t geometry.Triangle, pixels []byte, zBuffer []float32, mtl material.Material, lights []*render.Light, normal geometry.Vector3) {
	vertices := []geometry.Vector3{t.A, t.B, t.C}
	points := make([]geometry.Point, 3)
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

	rBase, gBase, bBase, _ := mtl.Color.RGBA()

	var width, height float32
	if mtl.MaterialType == material.Texture {
		width = float32(mtl.Image.Bounds().Max.X)
		height = float32(mtl.Image.Bounds().Max.Y)
	}
	
	// fmt.Println(r0, g0, b0, r1, g1, b1, r2, g2, b2)


	dist0 := v0.Distance(geometry.ZeroVector())
	dist1 := v1.Distance(geometry.ZeroVector())
	dist2 := v2.Distance(geometry.ZeroVector())
	
	wg := sync.WaitGroup{}

	for y := top; y < bottom; y++ {

		wg.Add(1)

		go func(y int) {
			var min int
			var max int
			if y < mid {
				a := f(p0, p1, y)
				b := f(p0, p2, y)
				min = utils.Min(a, b)
				max = utils.Max(a, b)
			} else {
				a := f(p1, p2, y)
				b := f(p0, p2, y)
				min = utils.Min(a, b)
				max = utils.Max(a, b)
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

				xCur := float64(weight0) * v0.X + float64(weight1) * v1.X + float64(weight2) * v2.X
				yCur := float64(weight0) * v0.Y + float64(weight1) * v1.Y + float64(weight2) * v2.Y
				zCur := float64(weight0) * v0.Z + float64(weight1) * v1.Z + float64(weight2) * v2.Z
				current := geometry.NewVector(xCur, yCur, zCur)

				var r, g, b float32

				//Iterate through all lights
				for _, light := range lights {
					if light.LightType == render.Ambient {
						rLight, gLight, bLight, _ := light.Color.RGBA()
						r += float32(rLight / 257)
						g += float32(gLight / 257)
						b += float32(bLight / 257)
					} else {
						// fmt.Println(light.ZBuffer[y*utils.RESOLUTION_X+x], current.Distance(light.Position))
						if light.ZBuffer[y*utils.RESOLUTION_X+x] >= current.Distance(light.Position) {
							percent := light.LightPercent(current, normal)
							// fmt.Println(percent)
							rLight, gLight, bLight, _ := light.Color.RGBA()
							r += float32(rLight / 257) * percent
							g += float32(gLight / 257) * percent
							b += float32(bLight / 257) * percent
						}
					}
					
				}

				z := weight0*dist0 + weight1*dist1 + weight2*dist2
				if x >= 0 && x < utils.RESOLUTION_X && y >= 0 && y < utils.RESOLUTION_Y {
					if z < zBuffer[y*utils.RESOLUTION_X+x] || zBuffer[y*utils.RESOLUTION_X+x] < 0 {
						zBuffer[y*utils.RESOLUTION_X+x] = z

						if mtl.MaterialType == material.Texture {
							u := weight0*v0.U + weight1*v1.U + weight2*v2.U
							v := weight0*v0.V + weight1*v1.V + weight2*v2.V
							u *= width
							v *= height
							rBase, gBase, bBase, _ = mtl.Image.At(int(u), int(v)).RGBA()
						}

						// fmt.Println(r, g, b)
						// fmt.Println(rBase, gBase, bBase)
						// fmt.Println("-----------------")

						pixels[(x+y*utils.RESOLUTION_X)*4+0] = uint8(r * float32(rBase) / 65535)
						pixels[(x+y*utils.RESOLUTION_X)*4+1] = uint8(g * float32(gBase) / 65535)
						pixels[(x+y*utils.RESOLUTION_X)*4+2] = uint8(b * float32(bBase) / 65535)
						pixels[(x+y*utils.RESOLUTION_X)*4+3] = 255
					}
				}
			}
			wg.Done()
		}(y)
	}

	wg.Wait()

}

func f(start, end geometry.Point, y int) int {
	height := end.Y - start.Y
	if height == 0 {
		height = 1
	}
	return start.X + (end.X-start.X)*(y-start.Y)/height
}