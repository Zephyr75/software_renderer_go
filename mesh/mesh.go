package mesh

import (
	"image"
	"overdrive/geometry"
	"overdrive/material"
	"overdrive/render"
	"sort"
	// "fmt"
)

type Mesh struct {
	Triangles []geometry.Triangle
	Material material.Material
	Position geometry.Vector3
	Rotation geometry.Vector3
}

func (m Mesh) Draw(img *image.RGBA, cam *render.Camera, lights []render.Light) {
	for i := range m.Triangles {
		(&m.Triangles[i].A).ResetLightAmount()
		(&m.Triangles[i].B).ResetLightAmount()
		(&m.Triangles[i].C).ResetLightAmount()
	}

	for i := range m.Triangles {
		// fmt.Println("--------")
		// fmt.Println("Vector Position: ", &m.Triangles[i])
		cam.ApplyCamera(&m.Triangles[i])
	}

	for i := range m.Triangles {
		normal := m.Triangles[i].Normal()
		for _, l := range lights {
			l.ApplyLight(&m.Triangles[i].A, normal)
			l.ApplyLight(&m.Triangles[i].B, normal)
			l.ApplyLight(&m.Triangles[i].C, normal)
		}
	}

	//TODO: sort triangles by distance to camera

	sort.SliceStable(m.Triangles, func(i, j int) bool {
		avgI := m.Triangles[i].Average()
		avgJ := m.Triangles[j].Average()
		distI := avgI.Distance(cam.Position)
		distJ := avgJ.Distance(cam.Position)
		return distI > distJ
	})

	for i := range m.Triangles {
		m.Triangles[i].Draw(img)
	}
	
}


func (m *Mesh) Rotate(rotation geometry.Vector3) {
	negativePosition := geometry.VectorNew(-m.Position.X, -m.Position.Y, -m.Position.Z)

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
	// for _, t := range m.geometry.Triangles {
	// 	t.A.AddAssign(position)
	// 	t.B.AddAssign(position)
	// 	t.C.AddAssign(position)
	// }
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
	v0 := geometry.VectorNew(size.X / 2, size.Y / 2, size.Z / 2)
	v1 := geometry.VectorNew(size.X / 2, size.Y / 2, -size.Z / 2)
	v2 := geometry.VectorNew(size.X / 2, -size.Y / 2, size.Z / 2)
	v3 := geometry.VectorNew(size.X / 2, -size.Y / 2, -size.Z / 2)
	v4 := geometry.VectorNew(-size.X / 2, size.Y / 2, size.Z / 2)
	v5 := geometry.VectorNew(-size.X / 2, size.Y / 2, -size.Z / 2)
	v6 := geometry.VectorNew(-size.X / 2, -size.Y / 2, size.Z / 2)
	v7 := geometry.VectorNew(-size.X / 2, -size.Y / 2, -size.Z / 2)
	triangles := make([]geometry.Triangle, 12)
	triangles[0] = geometry.TriangleNew(v1, v3, v7)
	triangles[1] = geometry.TriangleNew(v7, v5, v1)
	triangles[2] = geometry.TriangleNew(v2, v0, v4)
	triangles[3] = geometry.TriangleNew(v2, v4, v6)
	triangles[4] = geometry.TriangleNew(v0, v3, v1)
	triangles[5] = geometry.TriangleNew(v0, v2, v3)
	triangles[6] = geometry.TriangleNew(v4, v5, v7)
	triangles[7] = geometry.TriangleNew(v7, v6, v4)
	triangles[8] = geometry.TriangleNew(v5, v0, v1)
	triangles[9] = geometry.TriangleNew(v5, v4, v0)
	triangles[10] = geometry.TriangleNew(v7, v3, v2)
	triangles[11] = geometry.TriangleNew(v2, v6, v7)

	result := Mesh{triangles, material.Material{}, geometry.VectorZero(), geometry.VectorZero()}
	result.Translate(position)
	result.Rotate(rotation)
	return result
}