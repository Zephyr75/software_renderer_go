package render


import (
	"overdrive/geometry"
)


type Camera struct {
	Position geometry.Vector3
	Rotation geometry.Vector3
}

func (c *Camera) applyCameraVertex(v *geometry.Vector3) {
	//v.Rotate(c.Rotation)
	v.SubAssign(c.Position)
}

func (c *Camera) ApplyCamera(t *geometry.Triangle) {
	c.applyCameraVertex(&t.A)
	c.applyCameraVertex(&t.B)
	c.applyCameraVertex(&t.C)
}