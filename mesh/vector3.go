package mesh

import (
	"image/color"
	"math"
	"overdrive/material"
	"overdrive/utilities"
)

type Vector3 struct {
	X float64
	Y float64
	Z float64
	LightAmount color.Color
	Material material.Material
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
		material.Material{},
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

func (v Vector3) Add(v2 Vector3) Vector3 {
	return Vector3{
		v.X + v2.X,
		v.Y + v2.Y,
		v.Z + v2.Z,
		color.Black,
		material.Material{},
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
		material.Material{},
	}
}

func (v *Vector3) SubAssign(v2 Vector3) {
	v.X -= v2.X
	v.Y -= v2.Y
	v.Z -= v2.Z
}

func (v *Vector3) ApplyPerspective() {
	z0 := utilities.Z0()
	v.X = v.X * z0 / (z0 + v.Z)
	v.Y = v.Y * z0 / (z0 + v.Z)
	v.Z = z0
}

func (v *Vector3) CenterScreen() {
	v.X += utilities.RESOLUTION_X / 2
	v.Y += utilities.RESOLUTION_Y / 2
}

func (v Vector3) Distance(v2 Vector3) float64 {
	return math.Sqrt(math.Pow(v.X - v2.X, 2) + math.Pow(v.Y - v2.Y, 2) + math.Pow(v.Z - v2.Z, 2))
}