package geometry

type Triangle struct {
	A Vector3
	B Vector3
	C Vector3
}

func NewTriangle(a, b, c Vector3) Triangle {
	return Triangle{a, b, c}
}

func (t Triangle) Normal() Vector3 {
	v1 := t.B.Sub(t.A)
	v1.Normalize()
	v2 := t.C.Sub(t.A)
	v2.Normalize()
	return v1.Cross(v2)
}


//TODO: remove if unused
func (t Triangle) Average() Vector3 {
	return t.A.Add(t.B).Add(t.C).Div(3)
}

