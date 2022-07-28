package rectangle

type Rect struct {
	Length, Bredth float32
	Area           float32
}

func (r *Rect) AreaOf() float32 {
	r.Area = r.Length * r.Bredth
	return r.Area
}
