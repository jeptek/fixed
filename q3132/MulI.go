package q3132

// MulI returns x*y, use this when y is an integer for speed improvements.
//
// Special cases (in order):
//
//	MulI(NaN, y) = NaN
//	MulI(0, y) = 0
//	MulI(x, 0) = 0
//	MulI(±Inf, y) = ±Inf
//	MulI(±0, x) = ±0
//	MulI(x, y) = ±Inf if result is larger in magnitude than MaxValue
func MulI[T INT](x FX, y T) FX {
	_y := FX(y)
	// special cases:
	switch {
	case x == NaN:
		return NaN
	case x == 0 || _y == 0:
		return Zero
	case x == Inf || x == InfNeq:
		return Inf * sgn(x) * sgn(FX(y))
	case x == Infs || x == InfsNeq:
		return Infs * sgn(x) * sgn(FX(y))
	}

	sign := sgn(x) * sgn(_y)
	res := x * _y
	if sgn(res) != sign { // test overflow
		return Inf * sign
	}
	return res
}
