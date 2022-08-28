package mesh

import (
	"image"
	"overdrive/draw"
	"github.com/StephaneBunel/bresenham"
	"image/color"
)

type Triangle struct {
	A Vector3
	B Vector3
	C Vector3
}

func (t *Triangle) Draw(image *image.RGBA) {
	vertices := []Vector3{t.A, t.B, t.C}
	points := make([]draw.Point, 3)
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

	//iterate from top to bottom
	for y := top; y < bottom; y++ {
		var min int
		var max int
		if y > mid {
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
		bresenham.DrawLine(image, min, y, max, y, t.A.LightAmount)

	}
			
	
}

func f(start, end draw.Point, y int) int {
	height := end.Y - start.Y
	if height == 0 {
		height = 1
	}
	return start.X + (end.X - start.X) * (y - start.Y) / height
}
