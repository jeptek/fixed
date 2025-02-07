package q3132

// Add returns x+y
//
// Special cases are (in order, the order of x and y does not matter):
//
//	Add(NaN, y) = NaN
//	Add(-Inf, +Inf) = NaN
//	Add(±Inf, y) = ±Inf
//	Add(±0, ±0) = ±0
//	Add(±0, ∓0) = 0
//	Add(±0, y) = y
//	Add(x, y) = +Inf for x+y > 2147483647.999999999
//	Add(x, y) = -Inf for x+y < -2147483647.999999999
func Add(x, y FX) FX {
	// special cases:
	switch {
	case x == NaN || y == NaN:
		return NaN
	case x == Inf || x == InfNeq:
		if y == -x { // opposite inf
			return NaN
		} else {
			return x
		}

	case y == Inf || y == InfNeq:
		return y

	case x == Infs || x == InfsNeq:
		switch {
		case y == -x: // opposite infs
			return 0
		case y == Zero:
			return x
		default:
			return y
		}

	case y == Infs || y == InfsNeq:
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
