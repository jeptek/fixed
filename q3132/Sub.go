package q3132

// Add returns x-y
//
// Special cases are:
//
//	Sub(NaN, y) = NaN
//	Sub(x, NaN) = NaN
//	Sub(+Inf, +Inf) = NaN
//	Sub(-Inf, -Inf) = NaN
//	Sub(±Inf, y) = ±Inf
//	Sub(±0, ∓0) = ±0
//	Sub(±0, ±0) = 0
//	Sub(x, ±Inf) = ∓Inf
//	Sub(±0, y) = -y
//	Sub(x, ±0) = x
//	Sub(x, y) = +Inf for x-y > 2147483647.999999999
//	Sub(x, y) = -Inf for x-y < -2147483647.999999999
func Sub(x, y FX) FX {
	if y == NaN {
		return NaN
	}
	return Add(x, -y)
}
