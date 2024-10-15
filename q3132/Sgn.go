package q3132

// Sgn returns sign of x, apply it to FX number by using raw multiplication.
//
// Special cases are:
//
//	Sgn(NaN) = NaN
//	Sgn(0) = 0
//	Sgn(x>0) = 1
//	Sgn(x<0) = -1
func Sgn(x FX) FX {
	if x == NaN {
		return NaN
	} else if x == 0 {
		return Zero
	} else if x > 0 {
		return One
	} else {
		return -One
	}
}

// sgn returns sign of x, does not check for NaN
//
//	sgn(0) = 0
//	sgn(x>0) = raw(1)
//	sgn(x<0) = raw(-1)
func sgn(x FX) FX {
	if x == 0 {
		return Zero
	} else if x > 0 {
		return 1
	} else {
		return -1
	}
}
