package q3132

// Exp returns e**x, the base-e exponential of x.
//
// Special cases are:
//
//	Exp(NaN) = NaN
//	Exp(-Inf) = 0
//	Exp(0) = 1
//	Exp(±0) = 1
//	Exp(x<-22.180709777) = +0
//	Exp(x>21.487562591) = +Inf
func Exp(x FX) FX {
	const (
		// 21.487562591
		Overflow = 92288378601
		// -22.180709777
		Underflow = -95265423097
	)

	// special cases
	switch {
	case x == NaN:
		return NaN
	case x == -Inf:
		return 0
	case x == -Infs:
		return One
	case x == Infs:
		return One
	case x == Zero:
		return One
	case x >= Overflow:
		return Inf
	case x <= Underflow:
		return Infs
	case x == One:
		return E
	case x == -One:
		return Div(One, E)
	}

	if x < 0 {
		return Div(One, expRangeReduction(-x))
	}
	return expRangeReduction(x)
}

// Polynomial approximation for e^x using the first few terms of the Taylor series.
func expTaylorSeries(x FX) FX {
	// Start with the last coefficient.
	result := FX(106522)                     // 1/8!
	result = 852176 + fastMul(x, result)     // 1/7!
	result = 5965232 + fastMul(x, result)    // 1/6!
	result = 35791394 + fastMul(x, result)   // 1/5!
	result = 178956971 + fastMul(x, result)  // 1/4!
	result = 715827883 + fastMul(x, result)  // 1/3!
	result = 2147483648 + fastMul(x, result) // 1/2!
	result = ONE + fastMul(x, result)        // 1/1!
	result = ONE + fastMul(x, result)        // 1/0!
	return result
}

// Range reduction for exp(x): reduce x using e^x = 2^n * e^(x - n*ln(2))
func expRangeReduction(x FX) FX {
	// Find n such that x ≈ n * ln(2)
	n := x / Ln2

	// x = x - n * ln(2)
	xReduced := x - n*Ln2

	// Approximate e^x for the reduced x using Taylor series.
	expApprox := expTaylorSeries(xReduced)

	// Multiply result by 2^n
	if n > 0 {
		expApprox <<= uint64(n)
	} else {
		expApprox >>= uint64(-n)
	}

	return expApprox
}
