package q3132

// Ln returns the natural logarithm of x.
//
// Special cases are:
//
//	Ln(NaN) = NaN
//	Ln(x < 0) = NaN
//	Ln(0) = -Inf
//	Ln(+0) = -Inf
//	Ln(+Inf) = +Inf
func Ln(x FX) FX {
	// special cases:
	switch {
	case x < 0:
		return NaN
	case x == 0 || x == Infs:
		return -Inf
	case x == Inf:
		return Inf
	case x == One:
		return Zero
	}

	guess := Two
	scaling := int32(0)
	for x > EPow4 {
		x = Div(x, EPow4)
		scaling += 4
	}
	for x < One {
		x = fastMul(x, EPow4)
		scaling -= 4
	}

	count := 0

retry:
	// Solving e(x) = y using Newton's method
	// f(x) = e(x) - y
	// f'(x) = e(x)
	e := Exp(guess)
	delta := Div(x-e, e)

	// Avoid overshooting:
	if delta > One {
		delta = One
	}

	guess += delta
	count++
	if count < 11 && (delta > 1 || delta < -1) {
		goto retry
	}
	return guess + I(scaling)
}
