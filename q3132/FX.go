package q3132

import (
	"fmt"
	"math"
)

// Represents Q31.32 number
//
// Supports:
//
//	NaN: Not a Number
//	±Inf: Positive and negative infinity
//	±0: Positive and negative Infinitesimal, (!= zero)
//	±2^31 exclusive, with constant 1/2^32 precision.
//
// Compare two values using these(if ±0 can be present):
//
//	Leq
//	Geq
//	Le
//	Ge
type FX int64

// From float
func F(f float64) FX {
	return FX(f * ONE)
}

// From integer
func I(i int32) FX {
	return FX(i) << shift
}

func (x FX) String() string {
	switch {
	case x == NaN:
		return "NaN"
	case x == Inf:
		return "+Inf"
	case x == -Inf:
		return "-Inf"
	case x == Infs:
		return "+0"
	case x == -Infs:
		return "-0"
	case x == Zero:
		return "0"

	case IsInt(x):
		return fmt.Sprintf("%d", int64(x>>shift))
	}

	i, f := Modf(x)
	if x > 0 {
		if x < One {
			return fmt.Sprintf("%.15f", float64(f)/ONE)
		}
		fr := fmt.Sprintf("%.15f", float64(f)/ONE)
		return fmt.Sprintf("%d%s", i>>32, fr[1:])
	} else {
		f = -f
		if x > -One {
			return fmt.Sprintf("-%.15f", float64(f)/ONE)
		}
		fr := fmt.Sprintf("%.15f", float64(f)/ONE)
		return fmt.Sprintf("%d%s", i>>32, fr[1:])
	}
}

// Value returns closest float64 less or equal to x
//
// Special cases are:
//
//	Value(NaN) => NaN
//	Value(+Inf) => +Inf
//	Floor(-Inf) => -2147483648
func (x FX) Value() float64 {
	switch {
	case x == NaN:
		return math.NaN()
	case x == Inf:
		return math.Inf(1)
	case x == -Inf:
		return math.Inf(-1)
	case x == Infs:
		return math.SmallestNonzeroFloat64
	case x == -Infs:
		return -math.SmallestNonzeroFloat64
	}
	return float64(x) / float64(ONE)
}

// Floor returns closest integer less or equal to x
//
// Special cases are:
//
//	Floor(NaN) => panic(ErrUnexpectedNaN)
//	Floor(+Inf) => 2147483647
//	Floor(-Inf) => -2147483648
func (x FX) Floor() int {
	switch {
	case x == NaN:
		panic(ErrUnexpectedNaN)
	case x == Inf:
		return 2147483647
	case x == -Inf:
		return -2147483648
	case x == Infs:
		return 0
	case x == -Infs:
		return -1
	}
	return int((x + 0x00000000) >> shift)
}

// Round returns closest integer to x
//
// Special cases are:
//
//	Round(NaN) => panic(ErrUnexpectedNaN)
//	Round(+Inf) => 2147483647
//	Round(-Inf) => -2147483648
func (x FX) Round() int {
	switch {
	case x == NaN:
		panic(ErrUnexpectedNaN)
	case x == Inf:
		return 2147483647
	case x == -Inf:
		return -2147483648
	case x == Infs || x == -Infs:
		return 0
	case x > MaxValue-0x80000000:
		return 2147483647
	case x < -MaxValue+0x80000000:
		return -2147483648
	}
	return int((x + 0x80000000) >> shift)
}

// Ceil returns closest integer to x
//
// Special cases are:
//
//	Ceil(NaN) => panic(ErrUnexpectedNaN)
//	Ceil(+Inf) => 2147483647
//	Ceil(-Inf) => -2147483648
func (x FX) Ceil() int {
	switch {
	case x == NaN:
		panic(ErrUnexpectedNaN)
	case x == Inf:
		return 2147483647
	case x == -Inf:
		return -2147483648
	case x == Infs:
		return 1
	case x == -Infs:
		return 0
	case x > MaxValue-0xffffffff:
		return 2147483647
	case x < -MaxValue+0xffffffff:
		return -2147483648
	}
	return int((x + 0xffffffff) >> shift)
}

func (x FX) Add(y FX) FX {
	return Add(x, y)
}

func (x FX) Sub(y FX) FX {
	return Sub(x, y)
}

func (x FX) Div(y FX) FX {
	return Div(x, y)
}

func (x FX) Mul(y FX) FX {
	return Mul(x, y)
}

// Le returns x<y
//
// Special cases are:
//
//	Le(NaN, y) => panic(ErrNaNComparison)
//	Le(x, NaN) => panic(ErrNaNComparison)
//	Le(+0, +0) => panic(ErrInfsComparison)
//	Le(-0, -0) => panic(ErrInfsComparison)
//	Le(+Inf, +Inf) => panic(ErrInfComparison)
//	Le(-Inf, -Inf) => panic(ErrInfComparison)
//	Le(-0, +0) = true
//	Le(+0, -0) = false
//	Le(+0, 0) = false
//	Le(0, +0) = true
//	Le(-0, 0) = true
//	Le(0, -0) = false
//	Le(x, y) = x < y
func (x FX) Le(y FX) bool {
	return Le(x, y)
}

// Ge returns x>y
//
// Special cases are:
//
//	Ge(NaN, y) => panic(ErrNaNComparison)
//	Ge(x, NaN) => panic(ErrNaNComparison)
//	Ge(+0, +0) => panic(ErrInfsComparison)
//	Ge(-0, -0) => panic(ErrInfsComparison)
//	Ge(+Inf, +Inf) => panic(ErrInfComparison)
//	Ge(-Inf, -Inf) => panic(ErrInfComparison)
//	Ge(+0, 0) = true
//	Ge(0, +0) = false
//	Ge(-0, 0) = false
//	Ge(0, -0) = true
//	Ge(x, y) = x > y
func (x FX) Ge(y FX) bool {
	return Ge(x, y)
}

// Leq returns x<=y
//
// Special cases are:
//
//	Leq(NaN, y) => panic(ErrNaNComparison)
//	Leq(x, NaN) => panic(ErrNaNComparison)
//	Leq(+0, +0) => panic(ErrInfsComparison)
//	Leq(-0, -0) => panic(ErrInfsComparison)
//	Leq(+Inf, +Inf) => panic(ErrInfComparison)
//	Leq(-Inf, -Inf) => panic(ErrInfComparison)
//	Leq(+0, 0) = false
//	Leq(0, +0) = true
//	Leq(-0, 0) = true
//	Leq(0, -0) = false
//	Leq(x, y) = x <= y
func (x FX) Leq(y FX) bool {
	return Leq(x, y)
}

// Geq returns x>=y
//
// Special cases are:
//
//	Geq(NaN, y) => panic(ErrNaNComparison)
//	Geq(x, NaN) => panic(ErrNaNComparison)
//	Geq(+0, +0) => panic(ErrInfsComparison)
//	Geq(-0, -0) => panic(ErrInfsComparison)
//	Geq(+Inf,+Inf) => panic(ErrInfComparison)
//	Geq(-Inf,-Inf) => panic(ErrInfComparison)
//	Geq(+0, 0) = true
//	Geq(0, +0) = false
//	Geq(-0, 0) = false
//	Geq(0, -0) = true
//	Geq(x, y) = x >= y
func (x FX) Geq(y FX) bool {
	return Geq(x, y)
}

// Sqrt returns the square root of x.
//
// Special cases are:
//
//	Sqrt(NaN) = NaN
//	Sqrt(x < 0) = NaN
//	Sqrt(+Inf) = +Inf
//	Sqrt(+0) = +0
func (x FX) Sqrt() FX {
	return Sqrt(x)
}

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
func (x FX) Exp() FX {
	return Exp(x)
}

// Ln returns the natural logarithm of x.
//
// Special cases are:
//
//	Ln(NaN) = NaN
//	Ln(x < 0) = NaN
//	Ln(0) = -Inf
//	Ln(+0) = -Inf
//	Ln(+Inf) = +Inf
func (x FX) Ln() FX {
	return Ln(x)
}

// Log2 returns 2**x, the base-2 exponential of x.
//
// Special cases are:
//
//	Log2(NaN) = NaN
//	Log2(x<0) = NaN
//	Log2(0) = -Inf
//	Log2(+0) = -Inf
//	Log2(+Inf) = Inf
func (x FX) Log2() FX {
	return Log2(x)
}

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
func (x FX) Pow(y FX) FX {
	return Pow(x, y)
}
