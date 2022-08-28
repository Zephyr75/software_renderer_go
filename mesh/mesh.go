package mesh

import (
	"image"
	"image/color"
	"overdrive/material"
	"overdrive/geometry"
	"overdrive/render"
)

type Mesh struct {
	Triangles []geometry.Triangle
	Material material.Material
	Position geometry.Vector3
	Rotation geometry.Vector3
}

func (m *Mesh) Draw(image *image.RGBA, cam *render.Camera, lights []render.Light) {
	
	for _, t := range m.Triangles {
		cam.ApplyCamera(&t)
	}

	for _, t := range m.Triangles {
		normal := t.Normal()
		for _, l := range lights {
			l.ApplyLight(t.A, normal)
			l.ApplyLight(t.B, normal)
			l.ApplyLight(t.C, normal)
		}
	}

	//TODO sort triangles by distance to camera

	for _, t := range m.Triangles {
		t.Draw(image)
	}
	
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
	v0 := geometry.Vector3{size.X / 2, size.Y / 2, size.Z / 2, color.Black}
	v1 := geometry.Vector3{size.X / 2, size.Y / 2, -size.Z / 2, color.Black}
	v2 := geometry.Vector3{size.X / 2, -size.Y / 2, size.Z / 2, color.Black}
	v3 := geometry.Vector3{size.X / 2, -size.Y / 2, -size.Z / 2, color.Black}
	v4 := geometry.Vector3{-size.X / 2, size.Y / 2, size.Z / 2, color.Black}
	v5 := geometry.Vector3{-size.X / 2, size.Y / 2, -size.Z / 2, color.Black}
	v6 := geometry.Vector3{-size.X / 2, -size.Y / 2, size.Z / 2, color.Black}
	v7 := geometry.Vector3{-size.X / 2, -size.Y / 2, -size.Z / 2, color.Black}
	triangles := make([]geometry.Triangle, 8)
	triangles[0] = geometry.Triangle{v1, v3, v7}
	triangles[1] = geometry.Triangle{v7, v5, v1}
	triangles[2] = geometry.Triangle{v2, v0, v4}
	triangles[3] = geometry.Triangle{v2, v4, v6}
	triangles[4] = geometry.Triangle{v0, v3, v1}
	triangles[5] = geometry.Triangle{v0, v2, v3}
	triangles[6] = geometry.Triangle{v4, v5, v7}
	triangles[7] = geometry.Triangle{v7, v6, v4}
	triangles[8] = geometry.Triangle{v5, v0, v1}
	triangles[9] = geometry.Triangle{v5, v4, v0}
	triangles[10] = geometry.Triangle{v7, v3, v2}
	triangles[11] = geometry.Triangle{v2, v6, v7}

	result := Mesh{triangles, material.Material{}, geometry.VectorZero(), geometry.VectorZero()}
	//result.Translate(position)
	return result
}