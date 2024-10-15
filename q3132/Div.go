package q3132

import (
	"github.com/jeptek/fixed/internal"
)

// Div returns x/y
//
// Special cases are (in order):
//
//	Div(NaN, y) = NaN
//	Div(x, NaN) = NaN
//	Div(x, 0) = NaN
//	Div(0, y) = 0
//	Div(±Inf, ±Inf) = NaN
//	Div(±Inf, y) = ±Inf
//	Div(x, ±Inf) = 0
//	Div(±0, ±0) = NaN
//	Div(±0, y) = ±0
//	Div(x, ±0) = ±Inf
//	Div(x, y) = ±0 if result is smaller in magnitude than Iota
//	Div(x, y) = ±Inf if result is larger in magnitude than MaxValue
func Div(x, y FX) FX {
	// special cases:
	switch {
	case IsNaN(x) || IsNaN(y):
		return NaN
	case y == 0:
		return NaN
	case x == 0:
		return Zero
	case IsInf(x):
		switch {
		case IsInf(y):
			return NaN
		default:
			return Inf * sgn(x) * sgn(y)
		}
	case IsInf(y):
		return 0
	case IsInfs(x):
		switch {
		case IsInfs(y):
			return NaN
		default:
			return Infs * sgn(x) * sgn(y)
		}
	case IsInfs(y):
		return Inf * sgn(x) * sgn(y)
	case x == y:
		return ONE
	case x == -y:
		return -ONE
	}
	_x := abs(x)
	_y := abs(y)
	res := FX(div(uint64(_x), uint64(_y)))
	sign := sgn(x) * sgn(y)
	if res == Zero {
		return sign * Infs
	} else if (_y < ONE) && (res < _x) {
		return sign * Inf
	} else {
		return sign * res
	}
}

// Idea taken from go: math/bits.Div64()
func div(x uint64, y uint64) uint64 {
	// Main idea is to shift 'x' 33 bits left and then use 128bit integer division.
	// After that shift it back and use the last bit for rounding.

	u1 := x >> 31
	u0 := (x & mask) << 33
	hi := u1 / y
	r := u1 % y
	var lo uint64
	if r == 0 {
		lo = u0 / y
	} else {
		s := 64 - internal.Msb(y)
		y <<= s
		yn1 := y >> 32
		yn0 := y & mask
		un32 := r<<s | u0>>(64-s)
		un10 := u0 << s
		un1 := un10 >> 32
		un0 := un10 & mask
		q1 := un32 / yn1
		rhat := un32 - q1*yn1
		for q1 >= ONE || q1*yn0 > ONE*rhat+un1 {
			q1--
			rhat += yn1
		}
		un21 := un32*ONE + un1 - q1*y
		q0 := un21 / yn1
		rhat = un21 - q0*yn1
		for q0 >= ONE || q0*yn0 > ONE*rhat+un0 {
			q0--
			rhat += yn1
			if rhat >= ONE {
				break
			}
		}
		lo = q1*ONE + q0
	}
	return (hi << 31) + lo>>1 + (lo & 0x1)
}
