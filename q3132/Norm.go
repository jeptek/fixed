package q3132

// Norm returns normalized value of x
//
// Special cases are (cases are in order):
//
//	Norm(+Inf) = MaxValue
//	Norm(-Inf) = -MaxValue
//	Norm(+0) = Iota
//	Norm(-0) = -Iota
//	Norm(x) = x
func Norm(x FX) FX {
	// special cases:
	switch {
	case x == Inf:
		return MaxValue
	case x == -Inf:
		return -MaxValue
	case x == Infs:
		return Iota
	case x == -Infs:
		return -Iota
	}
	return x
}
