package mesh

import (
	// "fmt"
	"image"
	"overdrive/src/geometry"
	"overdrive/src/material"
	"overdrive/src/render"

	// "sort"
	"sync"
)

type Mesh struct {
	Triangles []geometry.Triangle
	Position  geometry.Vector3
	Rotation  geometry.Vector3
	Material  material.Material
}

/**
 * Render pipeline drawing the mesh on screen
 */
func (m Mesh) Draw(img *image.RGBA, zBuffer []float32, cam render.Camera, lights []render.Light) {

	wg := sync.WaitGroup{}
	for i := range m.Triangles {
		wg.Add(1)
		go func(t geometry.Triangle) {

			//Reset light passes
			(t.A).ResetLightAmount()
			(t.B).ResetLightAmount()
			(t.C).ResetLightAmount()

			//Compute light passes using forward rendering
			normal := t.Normal()
			for _, l := range lights {
				l.ApplyLight(&(t.A), normal)
				l.ApplyLight(&(t.B), normal)
				l.ApplyLight(&(t.C), normal)
			}

			//Apply camera transform
			t.A.AddAssign(cam.Position.Neg())
			t.B.AddAssign(cam.Position.Neg())
			t.C.AddAssign(cam.Position.Neg())
			t.A.Rotate(cam.Rotation.Neg())
			t.B.Rotate(cam.Rotation.Neg())
			t.C.Rotate(cam.Rotation.Neg())

			//Back-face culling
			normal = t.Normal()
			if normal.Z < 0 {
				//Rasterization
				t.Draw(img, zBuffer, m.Material)
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

/**
 * Rotates each vertex of the mesh around the origin
 */
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

/**
 * Creates a cube mesh
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

	result := Mesh{triangles, geometry.ZeroVector(), geometry.ZeroVector(), material.WhiteMaterial()}
	result.Translate(position)
	result.Rotate(rotation)
	return result
}

/**
 * Creates a mesh from a list of triangles
 */
func NewMesh(triangles []geometry.Triangle, position geometry.Vector3, rotation geometry.Vector3, material material.Material) Mesh {
	return Mesh{triangles, position, rotation, material}
}
