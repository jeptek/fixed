package q3132

// Mul returns x*y
//
// Special cases are (in order, the order of x and y does not matter):
//
//	Mul(NaN, y) = NaN
//	Mul(0, y) = 0
//	Mul(±Inf, y) = ±Inf
//	Mul(±0, y) = ±0
//	Mul(x, y) = ±0 if result is smaller in magnitude than Iota
//	Mul(x, y) = ±Inf if result is larger in magnitude than MaxValue
func Mul(x, y FX) FX {
	// special cases:
	switch {
	case x == NaN || y == NaN:
		return NaN
	case x == Zero || y == Zero:
		return Zero
	case x == Inf || x == InfNeq || y == Inf || y == InfNeq:
		return Inf * sgn(x) * sgn(y)
	case x == Infs || x == InfsNeq || y == Infs || y == InfsNeq:
		return Infs * sgn(x) * sgn(y)
	}

	u1 := uint64(x >> SHIFT)
	u0 := uint64(x & MASK)
	v1 := uint64(y >> SHIFT)
	v0 := uint64(y & MASK)

	w0 := u0 * v0
	t := u1*v0 + (w0 >> SHIFT)
	w1 := t & MASK
	w2 := uint64(int64(t) >> SHIFT)
	w1 += u0 * v1

	lo := uint64(x) * uint64(y)
	hi := u1*v1 + w2 + uint64(int64(w1)>>SHIFT)

	_hi := int64(hi) // overflow test:
	if int64(int32(_hi)) != _hi {
		return Inf * sgn(x) * sgn(y)
	}
	raw := (hi << 32) | (lo >> 32)
	raw += (lo >> 31) & 1 // Round to nearest, instead of rounding down.
	if raw == 0 {         // underflow test:
		return Infs * sgn(x) * sgn(y)
	}
	return FX(raw)
}

func fastMul(x, y FX) FX {
	xlo := uint64(x & MASK)
	xhi := uint64(x >> SHIFT)
	ylo := uint64(y & MASK)
	yhi := uint64(y >> SHIFT)

	lolo := xlo * ylo
	lohi := xlo * yhi
	hilo := xhi * ylo
	hihi := xhi * yhi

	var loResult = lolo>>SHIFT + (2147483648&lolo)>>31
	var hiResult = hihi << SHIFT

	return FX(loResult + lohi + hilo + hiResult)
}
