package render


import (
	"overdrive/mesh"
)


type Camera struct {
	Position mesh.Vector3
	Rotation mesh.Vector3
}

func (c *Camera) applyCameraVertex(v *mesh.Vector3) {
	v.Rotate(c.Rotation)
	v.SubAssign(c.Position)
}

func (c *Camera) ApplyCamera()