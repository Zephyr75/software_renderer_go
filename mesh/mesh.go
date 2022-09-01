package mesh

import (
	// "fmt"
	"image"
	"overdrive/geometry"
	"overdrive/render"

	// "sort"
	"sync"

)

type Mesh struct {
	Triangles []geometry.Triangle
	Position  geometry.Vector3
	Rotation  geometry.Vector3
}

func (m Mesh) Draw(img *image.RGBA, zBuffer []float32, cam render.Camera, lights []render.Light) {

	wg := sync.WaitGroup{}
	for i := range m.Triangles {
		wg.Add(1)
		go func(t geometry.Triangle) {
			(t.A).ResetLightAmount()
			(t.B).ResetLightAmount()
			(t.C).ResetLightAmount()
			t.A.AddAssign(cam.Position.Neg())
			t.B.AddAssign(cam.Position.Neg())
			t.C.AddAssign(cam.Position.Neg())
			t.A.Rotate(cam.Rotation.Neg())
			t.B.Rotate(cam.Rotation.Neg())
			t.C.Rotate(cam.Rotation.Neg())
			normal := t.Normal()
			for _, l := range lights {
				l.ApplyLight(&(t.A), normal)
				l.ApplyLight(&(t.B), normal)
				l.ApplyLight(&(t.C), normal)
			}
			if normal.Z < 0 {
				t.Draw(img, zBuffer)
			}
			wg.Done()
		}(m.Triangles[i])
	}
	wg.Wait()

}

func (m *Mesh) Rotate(rotation geometry.Vector3) {
	// m.Translate(m.Position.Neg())
	
	m.translateNoAssign(m.Position.Neg())
	m.rotateOrigin(rotation)
	m.translateNoAssign(m.Position)
	m.Rotation.AddAssign(rotation)
}

func (m *Mesh) rotateOrigin(rotation geometry.Vector3) {
	for i := range m.Triangles {
		m.Triangles[i].A.Rotate(rotation)
		m.Triangles[i].B.Rotate(rotation)
		m.Triangles[i].C.Rotate(rotation)
	}
}

func (m *Mesh) Translate(position geometry.Vector3) {
	m.translateNoAssign(position)
	m.Position.AddAssign(position)
}

func (m *Mesh) translateNoAssign(position geometry.Vector3) {
	for i := range m.Triangles {
		m.Triangles[i].A.AddAssign(position)
		m.Triangles[i].B.AddAssign(position)
		m.Triangles[i].C.AddAssign(position)
	}
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

	result := Mesh{triangles, geometry.ZeroVector(), geometry.ZeroVector()}
	result.Translate(position)
	result.Rotate(rotation)
	return result
}

func NewMesh(triangles []geometry.Triangle, position geometry.Vector3, rotation geometry.Vector3) Mesh {
	return Mesh{triangles, position, rotation}
}