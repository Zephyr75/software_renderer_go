package geometry

import (
	"image"
	"image/color"

	"sync"
	// "fmt"
)

type Triangle struct {
	A Vector3
	B Vector3
	C Vector3
}

func TriangleNew(a, b, c Vector3) Triangle {
	return Triangle{a, b, c}
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

func (t *Triangle) Draw(img *image.RGBA) {

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

	var p1, p2, p3 = points[0], points[1], points[2]
	var top, mid, bottom = p1.Y, p2.Y, p3.Y

	wg := sync.WaitGroup{}

	for y := top; y < bottom; y++ {

		wg.Add(1)

		go func(y int) {
			var min int
			var max int
			if y < mid {
				a := f(p1, p2, y)
				b := f(p1, p3, y)
				if a < b {
					min = a
					max = b
				} else {
					min = b
					max = a
				}
			} else {
				a := f(p2, p3, y)
				b := f(p1, p3, y)
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

			draw_line(int32(min), int32(y), int32(max), int32(y), t.A.LightAmount, img)

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

func draw_line(start_x int32, start_y int32, end_x int32, end_y int32, color color.Color, img *image.RGBA) {

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
