package q3132

// Pow returns x**y, the base-x exponential of y.
//
// Special cases are (in order):
//
//	Pow(NaN, y) = NaN
//	Pow(x, NaN) = NaN
//	Pow(x, ±0) = 1 for any x
//	Pow(1, y) = 1 for any y
//	Pow(x, 1) = x for any x
//	Pow(0, 0) = NaN
//	Pow(0, -Inf) = NaN
//	Pow(0, +Inf) = +Inf
//	Pow(0, x) = 1
//	Pow(-0, y ≤ +0) = NaN
//	Pow(-0, +Inf) = 0
//	Pow(-0, y) = NaN if y is not an integer
//	Pow(-0, y) = ±0  if y is an integer
//	Pow(+0, -Inf) = +Inf
//	Pow(+0, +Inf) = 0
//	Pow(+0, y<0) = +Inf
//	Pow(+0, y>0) = +0
//	Pow(±Inf, 0) = NaN
//	Pow(±Inf, y<0) = 0
//	Pow(-Inf, y>0) = -Inf
//	Pow(+Inf, y>0) = +Inf
//	Pow(x, -Inf) = 0
//	Pow(x, +Inf) = NaN for x < 0
//	Pow(x, +Inf) = +Inf for x > 1
//	Pow(x, +Inf) = 0 for x < 1
//	Pow(x, y) = NaN for x < 0 if y is not an integer
func Pow(x, y FX) FX {
	// special cases:
	switch {
	case x == NaN || y == NaN:
		return NaN
	case x == One:
		return One
	case x == Zero:
		switch {
		case y == 0 || y == -Inf:
			return NaN
		default:
			return Zero
		}
	case x == -Infs:
		switch {
		case y == 0:
			return One
		case y == Infs || y == -Infs:
			return NaN
		case y == Inf:
			return Zero
		case !IsInt(y):
			return NaN
		default:
			// y is positive integer
			i := abs(y) & 0xfffffff00000000
			if i%Two == One {
				return -Infs
			} else {
				return Infs
			}
		}
	case x == Infs:
		switch {
		case y == Inf:
			return Zero
		case y == 0 || y == Infs || y == -Infs:
			return One
		case y <= -One:
			return Inf
		case y < One:
			return One
		default: // y >= 1:
			return Infs
		}
	case x == -Inf:
		switch {
		case y == 0 || y == Infs || y == -Infs:
			return NaN
		case y < 0:
			return Zero
		default:
			return -Inf
		}
	case x == Inf:
		switch {
		case y == 0 || y == Infs || y == -Infs:
			return NaN
		case y < 0:
			return 0
		default:
			return Inf
		}
	}
	switch {
	case y == 0 || y == Infs || y == -Infs:
		return One
	case y == One:
		return x
	case y == -Inf:
		return Zero
	case y == Inf:
		switch {
		case x < 0:
			return NaN
		case x < One || x == Infs:
			return Zero
		default:
			return Inf
		}
	}

	sign := FX(1)
	if x < 0 { // determine sign of result:
		i, f := Modf(y)
		if f != 0 {
			return NaN
		}
		r := i % Two
		if r == One {
			sign = -1
		} else if r == -One {
			sign = -1
		}
		x = -x
	}

	log2 := Log2(x)
	if y < 0 {
		ip := Pow2(fastMul(-y, log2))
		if IsInf(ip) {
			return sign * Infs
		}
		return sign * Div(One, ip)
	}
	return sign * Pow2(fastMul(y, log2))
}
