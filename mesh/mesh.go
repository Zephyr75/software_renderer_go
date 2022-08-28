package mesh

import (
	"image/color"
	"overdrive/material"
)

type Mesh struct {
	Triangles []Triangle
	Material material.Material
	Position Vector3
	Rotation Vector3
}

func Cube(position Vector3, rotation Vector3, size Vector3) Mesh {
	v0 := Vector3{size.X / 2, size.Y / 2, size.Z / 2, color.Black}
	v1 := Vector3{size.X / 2, size.Y / 2, -size.Z / 2, color.Black}
	v2 := Vector3{size.X / 2, -size.Y / 2, size.Z / 2, color.Black}
	v3 := Vector3{size.X / 2, -size.Y / 2, -size.Z / 2, color.Black}
	v4 := Vector3{-size.X / 2, size.Y / 2, size.Z / 2, color.Black}
	v5 := Vector3{-size.X / 2, size.Y / 2, -size.Z / 2, color.Black}
	v6 := Vector3{-size.X / 2, -size.Y / 2, size.Z / 2, color.Black}
	v7 := Vector3{-size.X / 2, -size.Y / 2, -size.Z / 2, color.Black}
	triangles := make([]Triangle, 8)
	triangles[0] = Triangle{v1, v3, v7}
	triangles[1] = Triangle{v7, v5, v1}
	triangles[2] = Triangle{v2, v0, v4}
	triangles[3] = Triangle{v2, v4, v6}

	triangles[4] = Triangle{v0, v3, v1}
	triangles[5] = Triangle{v0, v2, v3}
	triangles[6] = Triangle{v4, v5, v7}
	triangles[7] = Triangle{v7, v6, v4}

	triangles[8] = Triangle{v5, v0, v1}
	triangles[9] = Triangle{v5, v4, v0}
	triangles[10] = Triangle{v7, v3, v2}
	triangles[11] = Triangle{v2, v6, v7}

	result := Mesh{triangles, material.Material{}, VectorZero(), VectorZero()}
	//result.Translate(position)
	return result
}



func (m *Mesh) Translate(position Vector3) {
	for i := range m.Triangles {
		m.Triangles[i].A.AddAssign(position)
		m.Triangles[i].B.AddAssign(position)
		m.Triangles[i].C.AddAssign(position)
	}
	// for _, t := range m.Triangles {
	// 	t.A.AddAssign(position)
	// 	t.B.AddAssign(position)
	// 	t.C.AddAssign(position)
	// }
	m.Position.AddAssign(position)
}

