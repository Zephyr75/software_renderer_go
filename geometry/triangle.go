package geometry

import (
	"image"
	"image/color"
	"overdrive/material"
	"overdrive/utilities"

	"sync"
)

type Triangle struct {
	A Vector3
	B Vector3
	C Vector3
	Material material.Material
}

func NewTriangle(a, b, c Vector3) Triangle {
	return Triangle{a, b, c, material.NewMaterial()}
}

func (t Triangle) Normal() Vector3 {
	v1 := t.B.Sub(t.A)
	v1_norm := v1.Norm()
	v1.DivAssign(v1_norm)
	v2 := t.C.Sub(t.A)
	v2_norm := v2.Norm()
	v2.DivAssign(v2_norm)
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

			//fmt.Println("y:", y, "min:", min, "max:", max)

			// image.Set(min, y, color.White)
			// image.Set(max, y, color.White)

			for x := min; x < max; x++ {
				current := Point{x, y}

				num1 := float32((p1.Y-p2.Y)*(current.X-p2.X) + (p2.X-p1.X)*(current.Y-p2.Y))
				num2 := float32((p2.Y-p0.Y)*(current.X-p2.X) + (p0.X-p2.X)*(current.Y-p2.Y))

				denum1 := float32((p1.Y-p2.Y)*(p0.X-p2.X) + (p2.X-p1.X)*(p0.Y-p2.Y))
				if denum1 == 0 {
					denum1 = 1
				}
				denum2 := float32((p1.Y-p2.Y)*(p0.X-p2.X) + (p2.X-p1.X)*(p0.Y-p2.Y))
				if denum2 == 0 {
					denum2 = 1
				}

				weight0 := num1 / denum1

				weight1 := num2 / denum2

				// fmt.Println("num1:", num1, "num2:", num2 ,"div1:", denum1, "div2:", denum2, "weight0:", weight0, "weight1:", weight1)

				weight2 := 1 - weight0 - weight1

				rA, gA, bA, _ := t.A.LightAmount.RGBA()
				rB, gB, bB, _ := t.B.LightAmount.RGBA()
				rC, gC, bC, _ := t.C.LightAmount.RGBA()

				r := float32(weight0) * float32(rA) + float32(weight1)*float32(rB) + float32(weight2)*float32(rC)
				g := float32(weight0) * float32(gA) + float32(weight1)*float32(gB) + float32(weight2)*float32(gC)
				b := float32(weight0) * float32(bA) + float32(weight1)*float32(bB) + float32(weight2)*float32(bC)
				r = r / 257
				g = g / 257
				b = b / 257

				tR, tG, tB, _ := t.Material.Color.RGBA() 
				tR /= 257 * 255
				tG /= 257 * 255
				tB /= 257 * 255

				z := weight0 * t.A.Distance(ZeroVector()) + weight1 * t.B.Distance(ZeroVector()) + weight2 * t.C.Distance(ZeroVector())

				// fmt.Println("y: ", y, "x: ", x)

				if x >= 0 && x < utilities.RESOLUTION_X && y >= 0 && y < utilities.RESOLUTION_Y {
					// fmt.Println("z: ", z, "zBuffer[y * utilities.RESOLUTION_X + x]: ", zBuffer[y * utilities.RESOLUTION_X + x])
					if z < zBuffer[y * utilities.RESOLUTION_X + x] {
						zBuffer[y * utilities.RESOLUTION_X + x] = z
						img.Set(x, y, color.RGBA{uint8(r * float32(tR)), uint8(g * float32(tG)), uint8(b * float32(tB)), 255})
						// fmt.Println("done")
					}
				} 
				// fmt.Println(z)
				

				// img.Set(x, y, color.RGBA{uint8(r * float32(tR)), uint8(g * float32(tG)), uint8(b * float32(tB)), 255})
			}

			// drawLine(int32(min), int32(y), int32(max), int32(y), color.RGBA{uint8(r), uint8(g), uint8(b), 255}, img)

			// min = max
			// max = min
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

func drawLine(start_x int32, start_y int32, end_x int32, end_y int32, color color.Color, img *image.RGBA) {

	// Bresenham's
	var cx int32 = start_x
	var cy int32 = start_y

	var dx int32 = end_x - cx
	var dy int32 = end_y - cy
	if dx < 0 {
		dx = 0 - dx
	}
	if dy < 0 {
		dy = 0 - dy
	}

	var sx int32
	var sy int32
	if cx < end_x {
		sx = 1
	} else {
		sx = -1
	}
	if cy < end_y {
		sy = 1
	} else {
		sy = -1
	}
	var err int32 = dx - dy

	var n int32
	for n = 0; n < 1000; n++ {
		img.Set(int(cx), int(cy), color)
		//draw_point(cx, cy, color, screen)
		if (cx == end_x) && (cy == end_y) {
			return
		}
		var e2 int32 = 2 * err
		if e2 > (0 - dy) {
			err = err - dy
			cx = cx + sx
		}
		if e2 < dx {
			err = err + dx
			cy = cy + sy
		}
	}
}
