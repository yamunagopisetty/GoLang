package square

type Square struct {
	Side float32
	Area float32
}

func (s *Square) AreaOf() float32 {
	s.Area = s.Side * s.Side
	return s.Area
}
