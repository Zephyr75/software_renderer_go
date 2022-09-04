package render

import (
	"overdrive/src/geometry"
)

type Camera struct {
	Position geometry.Vector3
	Rotation geometry.Vector3
}

func (c Camera) ApplyCamera(t *geometry.Triangle) {
	t.A.SubAssign(c.Position)
	t.B.SubAssign(c.Position)
	t.C.SubAssign(c.Position)
}
