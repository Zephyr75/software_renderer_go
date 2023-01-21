package render

import (
	"overdrive/src/geometry"
)

type Camera struct {
	Position geometry.Vector3
	Direction geometry.Vector3
}

func (c Camera) ApplyCamera(t *geometry.Triangle) {
	t.A.SubAssign(c.Position)
	t.B.SubAssign(c.Position)
	t.C.SubAssign(c.Position)
	t.A.Rotate(c.Direction.Neg())
	t.B.Rotate(c.Direction.Neg())
	t.C.Rotate(c.Direction.Neg())
}

/*
 __   __        __  ___  __        __  ___  __   __   __  
/  ` /  \ |\ | /__`  |  |__) |  | /  `  |  /  \ |__) /__` 
\__, \__/ | \| .__/  |  |  \ \__/ \__,  |  \__/ |  \ .__/                                                    

*/

func NewCamera(position, direction geometry.Vector3) Camera {
	return Camera{position, direction}
}