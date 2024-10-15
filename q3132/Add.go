package q3132

// Add returns x+y
//
// Special cases are (in order, the order of x and y does not matter):
//
//	Add(NaN, y) = NaN
//	Add(-Inf, +Inf) = NaN
//	Add(±Inf, y) = ±Inf
//	Add(±0, y) = y
//	Add(x, y) = +Inf for x+y > 2147483647.999999999
//	Add(x, y) = -Inf for x+y < -2147483647.999999999
func Add(x, y FX) FX {
	// special cases:
	switch {
	case x == NaN || y == NaN:
		return NaN
	case IsInf(x):
		if y == -x { // opposite inf
			return NaN
		} else {
			return x
		}
	case IsInf(y):
		return y

	case IsInfs(x):
		switch {
		case y == -x: // opposite infs
			return NaN
		case y == Zero:
			return x
		default:
			return y
		}

	case IsInfs(y):
		switch {
		case x == Zero:
			return y
		default:
			return x
		}
	}

	// check overflow:
	if x > 0 && y > 0 {
		if x > MaxValue-y {
			return Inf
		}
	} else if x < 0 && y < 0 {
		if x < MinValue-y {
			return -Inf
		}
	}
	return x + y
}
