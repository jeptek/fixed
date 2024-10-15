package q3132

// Abs returns the absolute value of x.
//
// Special cases are:
//
//	Abs(NaN) = NaN
func Abs(x FX) FX {
	switch {
	case x == NaN:
		return NaN
	case x < 0:
		return -x
	default:
		return x
	}
}

// Will skip NaN check
func abs(x FX) FX {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
