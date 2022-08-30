package geometry

import (
	"image/color"
	"math"
	"overdrive/utilities"
)

type Vector3 struct {
	X float64
	Y float64
	Z float64
	LightAmount color.Color
}

func (v *Vector3) ResetLightAmount() {
	v.LightAmount = color.Black
}

func (v Vector3) Norm() float64 {
    return math.Sqrt(v.X * v.X + v.Y * v.Y + v.Z * v.Z)
}

func (v *Vector3) Normalize() {
	norm := v.Norm()
	v.X /= norm
	v.Y /= norm
	v.Z /= norm
}

func (v Vector3) Cross(v2 Vector3) Vector3 {
	return Vector3{
		v.Y * v2.Z - v.Z * v2.Y,
		v.Z * v2.X - v.X * v2.Z,
		v.X * v2.Y - v.Y * v2.X,
		color.Black,
	}
}

func (v Vector3) Dot(v2 Vector3) float64 {
	return v.X * v2.X + v.Y * v2.Y + v.Z * v2.Z
}

func (v *Vector3) Rotate(r Vector3) {
	var x, y, z, rx, ry, rz = v.X, v.Y, v.Z, r.X, r.Y, r.Z
	v.X = x * math.Cos(rz) * math.Cos(ry) + 
		  y * (math.Cos(rz) * math.Sin(ry) * math.Sin(rx) - math.Sin(rz) * math.Cos(rx)) + 
		  z * (math.Cos(rz) * math.Sin(ry) * math.Cos(rx) + math.Sin(rz) * math.Sin(rx))
	v.Y = x * math.Sin(rz) * math.Cos(ry) +
		  y * (math.Sin(rz) * math.Sin(ry) * math.Sin(rx) + math.Cos(rz) * math.Cos(rx)) +
		  z * (math.Sin(rz) * math.Sin(ry) * math.Cos(rx) - math.Cos(rz) * math.Sin(rx))
	v.Z = -x * math.Sin(ry) + y * math.Cos(ry) * math.Sin(rx) + z * math.Cos(ry) * math.Cos(rx)
}

func (v Vector3) Converted() Point {
	v.applyPerspective()
	v.centerScreen()
	return v.toPoint()
}

func (v *Vector3) applyPerspective() {
	z0 := utilities.Z0()
	v.X = v.X * z0 / (z0 + v.Z)
	v.Y = v.Y * z0 / (z0 + v.Z)
	v.Z = z0
}

func (v *Vector3) centerScreen() {
	v.X += utilities.RESOLUTION_X / 2
	v.Y += utilities.RESOLUTION_Y / 2
}

func (v Vector3) toPoint() Point {
	return Point{int(v.X), int(v.Y)}
}


func (v Vector3) Distance(v2 Vector3) float64 {
	return math.Sqrt(math.Pow(v.X - v2.X, 2) + math.Pow(v.Y - v2.Y, 2) + math.Pow(v.Z - v2.Z, 2))
}

/*
 ██████  ██████  ███    ██ ███████ ████████ ██████  ██    ██  ██████ ████████  ██████  ██████  ███████ 
██      ██    ██ ████   ██ ██         ██    ██   ██ ██    ██ ██         ██    ██    ██ ██   ██ ██      
██      ██    ██ ██ ██  ██ ███████    ██    ██████  ██    ██ ██         ██    ██    ██ ██████  ███████ 
██      ██    ██ ██  ██ ██      ██    ██    ██   ██ ██    ██ ██         ██    ██    ██ ██   ██      ██ 
 ██████  ██████  ██   ████ ███████    ██    ██   ██  ██████   ██████    ██     ██████  ██   ██ ███████ 
*/

func VectorZero() Vector3 {
	return Vector3{0, 0, 0, color.Black}
}

func VectorNew(x, y, z float64) Vector3 {
	return Vector3{x, y, z, color.Black}
}

/*
 ██████  ██████  ███████ ██████   █████  ████████  ██████  ██████  ███████ 
██    ██ ██   ██ ██      ██   ██ ██   ██    ██    ██    ██ ██   ██ ██      
██    ██ ██████  █████   ██████  ███████    ██    ██    ██ ██████  ███████ 
██    ██ ██      ██      ██   ██ ██   ██    ██    ██    ██ ██   ██      ██ 
 ██████  ██      ███████ ██   ██ ██   ██    ██     ██████  ██   ██ ███████                               
*/

func (v Vector3) Add(v2 Vector3) Vector3 {
	return Vector3{
		v.X + v2.X,
		v.Y + v2.Y,
		v.Z + v2.Z,
		color.Black,
	}
}

func (v *Vector3) AddAssign(v2 Vector3) {
	v.X += v2.X
	v.Y += v2.Y
	v.Z += v2.Z
}

func (v Vector3) Sub(v2 Vector3) Vector3 {
	return Vector3{
		v.X - v2.X,
		v.Y - v2.Y,
		v.Z - v2.Z,
		color.Black,
	}
}

func (v *Vector3) SubAssign(v2 Vector3) {
	v.X -= v2.X
	v.Y -= v2.Y
	v.Z -= v2.Z
}

func (v Vector3) Mul(x float64) Vector3 {
	return Vector3{
		v.X * x,
		v.Y * x,
		v.Z * x,
		color.Black,
	}
}

func (v *Vector3) MulAssign(x float64) {
	v.X *= x
	v.Y *= x
	v.Z *= x
}

func (v Vector3) Div(x float64) Vector3 {
	return Vector3{
		v.X / x,
		v.Y / x,
		v.Z / x,
		color.Black,
	}
}

func (v *Vector3) DivAssign(x float64) {
	v.X /= x
	v.Y /= x
	v.Z /= x
}

func (v Vector3) Neg() Vector3 {
	return Vector3{
		-v.X,
		-v.Y,
		-v.Z,
		color.Black,
	}
}

func (v *Vector3) NegAssign() {
	v.X = -v.X
	v.Y = -v.Y
	v.Z = -v.Z
}