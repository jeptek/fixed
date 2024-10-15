package q3132

import "github.com/jeptek/fixed/internal"

// Log2 returns the log with base 2.
//
// Special cases are:
//
//	Log2(NaN) = NaN
//	Log2(x<0) = NaN
//	Log2(0) = -Inf
//	Log2(+0) = -Inf
//	Log2(+Inf) = Inf
func Log2(x FX) FX {
	// special cases:
	switch {
	case x < 0:
		return NaN
	case x == 0 || x == Infs:
		return -Inf
	case x == Inf:
		return Inf
	}

	if (x & (x - 1)) == 0 { // has one bit => is power of 2
		m := internal.Msb(uint64(x))
		return FX((m - 33) << shift)
	}

	// If the input is less than one, the result is -log2(1.0 / in)
	if x < One {
		inverse := Div(One, x)
		return -innerLog2(inverse)
	}

	return innerLog2(x)
}

func innerLog2(x FX) FX {
	res := FX(0)
	for x >= Two {
		res++
		x = (x >> 1) + (x & 1)
	}
	for i := 32; i > 0; i-- {
		x = fastMul(x, x)
		res <<= 1
		if x >= Two {
			res |= 1
			x = (x >> 1) + (x & 1)
		}
	}
	return res
}
