package q3132

import (
	"github.com/jeptek/fixed/internal"
)

// Sqrt returns the square root of x.
//
// Special cases are:
//
//	Sqrt(NaN) = NaN
//	Sqrt(x < 0) = NaN
//	Sqrt(+Inf) = +Inf
//	Sqrt(+0) = +0
func Sqrt(x FX) FX {

	// special cases
	switch {
	case x < 0:
		return NaN
	case x == 0:
		return 0
	case x == Infs:
		return Infs
	case x == Inf:
		return Inf
	case x == One:
		return One
	}

	m := (internal.Msb(uint64(x)) - 32)
	res := uint64(ONE)
	if m > 0 {
		res <<= ((m + 1) / 2)
	} else {
		res >>= ((1 - m) / 2)
	}

	xu := uint64(x)
	res = (div(xu, res) + res) >> 1
	res = (div(xu, res) + res) >> 1
	res = (div(xu, res) + res) >> 1
	res = (div(xu, res) + res) >> 1
	return FX(res)
}
