package mesh

import (
	// "fmt"
	"image"
	"overdrive/geometry"
	"overdrive/material"
	"overdrive/render"
	"sort"
)

type Mesh struct {
	Triangles []geometry.Triangle
	Material  material.Material
	Position  geometry.Vector3
	Rotation  geometry.Vector3
}

func (m Mesh) Draw(img *image.RGBA, cam render.Camera, lights []render.Light) {

	triangles := make([]geometry.Triangle, 12)

	copy(triangles, m.Triangles[:])

	for i := range triangles {
		(triangles[i].A).ResetLightAmount()
		(triangles[i].B).ResetLightAmount()
		(triangles[i].C).ResetLightAmount()
	}

	for i := range triangles {
		cam.ApplyCamera(&(triangles[i]))
	}

	for i := range triangles {
		normal := triangles[i].Normal()
		for _, l := range lights {
			l.ApplyLight(&(triangles[i].A), normal)
			l.ApplyLight(&(triangles[i].B), normal)
			l.ApplyLight(&(triangles[i].C), normal)
		}
	}

	sort.SliceStable(triangles, func(i, j int) bool {
		avgI := triangles[i].Average()
		avgJ := triangles[j].Average()
		distI := avgI.Distance(cam.Position)
		distJ := avgJ.Distance(cam.Position)
		return distI > distJ
	})


	for i := range triangles {
		normal := triangles[i].Normal()
		if normal.Z < 0 {
			triangles[i].Draw(img)
		}
	}

}

func (m *Mesh) Rotate(rotation geometry.Vector3) {
	negativePosition := geometry.NewVector(-m.Position.X, -m.Position.Y, -m.Position.Z)

	m.Translate(negativePosition)
	m.rotateOrigin(rotation)
	m.Translate(m.Position)
	m.Rotation.AddAssign(rotation)
}

func (m *Mesh) rotateOrigin(rotation geometry.Vector3) {
	for i := range m.Triangles {
		m.Triangles[i].A.Rotate(rotation)
		m.Triangles[i].B.Rotate(rotation)
		m.Triangles[i].C.Rotate(rotation)
	}
	m.Position.AddAssign(rotation)
}

func (m *Mesh) Translate(position geometry.Vector3) {
	for i := range m.Triangles {
		m.Triangles[i].A.AddAssign(position)
		m.Triangles[i].B.AddAssign(position)
		m.Triangles[i].C.AddAssign(position)
	}
	m.Position.AddAssign(position)
}

/*
 ██████  ██████  ███    ██ ███████ ████████ ██████  ██    ██  ██████ ████████  ██████  ██████  ███████
██      ██    ██ ████   ██ ██         ██    ██   ██ ██    ██ ██         ██    ██    ██ ██   ██ ██
██      ██    ██ ██ ██  ██ ███████    ██    ██████  ██    ██ ██         ██    ██    ██ ██████  ███████
██      ██    ██ ██  ██ ██      ██    ██    ██   ██ ██    ██ ██         ██    ██    ██ ██   ██      ██
 ██████  ██████  ██   ████ ███████    ██    ██   ██  ██████   ██████    ██     ██████  ██   ██ ███████
*/

func Cube(position geometry.Vector3, rotation geometry.Vector3, size geometry.Vector3) Mesh {
	v0 := geometry.NewVector(size.X/2, size.Y/2, size.Z/2)
	v1 := geometry.NewVector(size.X/2, size.Y/2, -size.Z/2)
	v2 := geometry.NewVector(size.X/2, -size.Y/2, size.Z/2)
	v3 := geometry.NewVector(size.X/2, -size.Y/2, -size.Z/2)
	v4 := geometry.NewVector(-size.X/2, size.Y/2, size.Z/2)
	v5 := geometry.NewVector(-size.X/2, size.Y/2, -size.Z/2)
	v6 := geometry.NewVector(-size.X/2, -size.Y/2, size.Z/2)
	v7 := geometry.NewVector(-size.X/2, -size.Y/2, -size.Z/2)
	triangles := make([]geometry.Triangle, 12)
	//var triangles [12]geometry.Triangle
	triangles[0] = geometry.NewTriangle(v1, v3, v7)
	triangles[1] = geometry.NewTriangle(v7, v5, v1)
	triangles[2] = geometry.NewTriangle(v2, v0, v4)
	triangles[3] = geometry.NewTriangle(v2, v4, v6)
	triangles[4] = geometry.NewTriangle(v0, v3, v1)
	triangles[5] = geometry.NewTriangle(v0, v2, v3)
	triangles[6] = geometry.NewTriangle(v4, v5, v7)
	triangles[7] = geometry.NewTriangle(v7, v6, v4)
	triangles[8] = geometry.NewTriangle(v5, v0, v1)
	triangles[9] = geometry.NewTriangle(v5, v4, v0)
	triangles[10] = geometry.NewTriangle(v7, v3, v2)
	triangles[11] = geometry.NewTriangle(v2, v6, v7)

	result := Mesh{triangles, material.NewMaterial(), geometry.ZeroVector(), geometry.ZeroVector()}
	result.Translate(position)
	result.Rotate(rotation)
	return result
}
