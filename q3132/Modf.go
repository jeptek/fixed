package q3132

// Modf returns integer and fractional that sum to x. Both will have same sign as x.
//
// Special cases are:
//
//	Modf(NaN) = NaN, NaN
//	Modf(±Inf) = ±Inf, NaN
//	Modf(±0) = 0, ±Iota
func Modf(x FX) (i FX, f FX) {
	switch {
	case x == NaN:
		return NaN, NaN
	case x == Inf:
		return Inf, NaN
	case x == -Inf:
		return -Inf, NaN
	case x == Infs:
		return 0, Iota
	case x == -Infs:
		return 0, -Iota
	}
	sign := FX(1)
	if x < 0 {
		sign = -1
		x = -x
	}
	return sign * (x &^ MASK), sign * (x & MASK)
}
