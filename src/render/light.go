package render

import (
	"image/color"
	"overdrive/src/geometry"
	"overdrive/src/utils"
	"sync"
)

// Enum for light types
type LightType byte

const (
	Directional LightType = 0
	Point       LightType = 1
	Ambient     LightType = 2
)

// Light component in the scene of a given type and color
type Light struct {
	Position  geometry.Vector3
	Direction geometry.Vector3
	LightType LightType
	Color     color.Color
	Length    float64
	ZBuffer   []float32
}




// Compute how much light a vertex gets
func (l Light) LightPercent(v geometry.Vector3, normal geometry.Vector3) float32 {
	var percentToApply float64
	switch l.LightType {
		case Ambient:
			percentToApply = 1
		case Directional:
			l.Direction.Normalize()
			percentToApply = normal.Dot(l.Direction)
		case Point:
			direction := l.Position.Sub(v)
			dim := 1 - direction.Norm()/l.Length
			direction.Normalize()
			// fmt.Println(normal.Dot(direction), dim)
			percentToApply = normal.Dot(direction) * dim
	}
	if percentToApply < 0 {
		percentToApply = 0
	}
	return float32(percentToApply)
}


// Fills zBuffer with the depth of each pixel relative to the light
func (l Light) FillBuffer(t geometry.Triangle) {
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
	dist0 := v0.Distance(l.Position)
	dist1 := v1.Distance(l.Position)
	dist2 := v2.Distance(l.Position)

	var wg sync.WaitGroup
	wg.Add(bottom - top)
	for y := top; y < bottom; y++ {
		go func(y int) {
			defer wg.Done()
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
				} //TODO: fix this
				denom2 := (p1.Y-p2.Y)*(p0.X-p2.X) + (p2.X-p1.X)*(p0.Y-p2.Y)
				if denom2 == 0 {
					denom2 = 1
				} //TODO: fix this
				weight0 := float32(num1) / float32(denom1)
				weight1 := float32(num2) / float32(denom2)
				weight2 := 1 - weight0 - weight1

				z := weight0*dist0 + weight1*dist1 + weight2*dist2
				if x >= 0 && x < utils.RESOLUTION_X && y >= 0 && y < utils.RESOLUTION_Y {
					if z < l.ZBuffer[y*utils.RESOLUTION_X+x] || l.ZBuffer[y*utils.RESOLUTION_X+x] < 0 {
						l.ZBuffer[y*utils.RESOLUTION_X+x] = z
					}
				}
			}
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

/*
 ██████  ██████  ███    ██ ███████ ████████ ██████  ██    ██  ██████ ████████  ██████  ██████  ███████
██      ██    ██ ████   ██ ██         ██    ██   ██ ██    ██ ██         ██    ██    ██ ██   ██ ██
██      ██    ██ ██ ██  ██ ███████    ██    ██████  ██    ██ ██         ██    ██    ██ ██████  ███████
██      ██    ██ ██  ██ ██      ██    ██    ██   ██ ██    ██ ██         ██    ██    ██ ██   ██      ██
 ██████  ██████  ██   ████ ███████    ██    ██   ██  ██████   ██████    ██     ██████  ██   ██ ███████
*/

func AmbientLight(color color.Color) Light {
	return Light{geometry.ZeroVector(), geometry.ZeroVector(), Ambient, color, 0, nil}
}

func PointLight(position, direction geometry.Vector3, color color.Color, length float64) Light {
	zBuffer := make([]float32, utils.RESOLUTION_X*utils.RESOLUTION_Y)
	return Light{position, direction, Point, color, length, zBuffer}
}

func DirectionalLight(direction geometry.Vector3, color color.Color) Light {
	zBuffer := make([]float32, utils.RESOLUTION_X*utils.RESOLUTION_Y)
	return Light{geometry.ZeroVector(), direction, Directional, color, 0, zBuffer}
}
