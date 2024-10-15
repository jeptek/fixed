package q3132

// Pow2 returns 2**x
//
// Special cases are:
//
//	Pow2(NaN) = NaN
//	Pow2(±0) = +0
//	Pow2(±Inf) = +Inf
func Pow2(x FX) FX {
	const (
		Overflow  = 0x0000001F00000000
		Underflow = -0x0000002000000000
	)

	if x == NaN {
		return NaN
	} else if x == 0 || x == Infs || x == -Infs {
		return One
	} else if x == -Inf {
		return Zero
	} else if x >= Overflow {
		return Inf
	} else if x <= Underflow {
		return Infs
	}

	// Avoid negative arguments by exploiting that exp(-x) = 1/exp(x).
	neg := x < 0
	if neg {
		x = -x
	}

	/* The algorithm is based on the power series for exp(x):
	 * http://en.wikipedia.org/wiki/Exponential_function#Formal_definition
	 *
	 * From term n, we get term n+1 by multiplying with x/n.
	 * When the sum term drops to zero, we can stop summing.
	 */
	integer, frac := Modf(x)
	res := FX(One)
	term := FX(One)

	var i FX = 1
	for term > 1 {
		term = fastMul(term, fastMul(Ln2, Div(frac, i<<shift)))
		res += term
		i++
	}
	res = res << (integer >> shift)
	if neg {
		return Div(One, res)
	}

	return res
}
