package q3132

const (
	shift = 32
	mask  = 1<<shift - 1

	ONE = 1 << shift

	// Not a Number
	NaN FX = -9223372036854775808
	// Infinite, nonreal
	//
	// Defines all values bigger in magnitude than largest representable value, use -Inf for negative.
	Inf FX = 9223372036854775807

	// Infinite, nonreal
	//
	// Defines all values bigger in magnitude than largest representable value, use -Inf for negative.
	InfNeq FX = -Inf

	// Infinitesimal, non zero, nonreal
	//
	// Defines all values that are smaller in magnitude than Iota.
	Infs FX = 9223372036854775806

	// Infinitesimal, non zero, nonreal
	//
	// Defines all values that are smaller in magnitude than Iota.
	InfsNeq FX = -Infs

	// Smallest non zero real value, adjacent values differ always by Iota.
	//
	// Value:
	//	= 1/2^32
	//	≈ 2.3283064e-10
	Iota FX = 1

	// 2147483647.99999999930150806903839111328125
	MaxValue FX = 9223372036854775805
	// -2147483647.99999999930150806903839111328125
	MinValue FX = -9223372036854775805

	// 2.71828182846
	//
	// Euler's number
	E FX = 11674931554
	// 54.5981500331
	//
	// (Euler's number)^4
	EPow4 FX = 234497268814
	// 0.5
	Half FX = ONE >> 1
	// 0.69314718056
	//
	// ln(2)
	Ln2 FX = 2977044471
	// 1
	One FX = ONE
	// 3.14159265359
	//
	// Pi
	Pi FX = 13493037704
	// 1.57079632679
	//
	// Pi/2
	PiDiv2 FX = 6746518852
	// 0.31830988618
	//
	// 1/Pi
	PiInv FX = 1367130551
	// 6.28318530718
	//
	// 2*Pi
	PiMul2 FX = 26986075409
	// 1.5
	ThreeHalfs FX = 3 << (shift - 1)
	// 2
	Two FX = 2 << shift
	// 0
	Zero FX = 0
)

// Returns true when x is NaN
func IsNaN(x FX) bool {
	return x == NaN
}

// Returns true when x is -Inf or +Inf
func IsInf(x FX) bool {
	return x == Inf || x == -Inf
}

// Returns true when x is -0 or +0
func IsInfs(x FX) bool {
	return x == Infs || x == -Infs
}

// Returns true when x is not: NaN, ±Inf or ±0
func IsReal(x FX) bool {
	switch {
	case x == NaN:
		return false
	case x == Infs:
		return false
	case x == -Infs:
		return false
	case x == Inf:
		return false
	case x == -Inf:
		return false
	}
	return true
}

// Returns true if x is an integer
func IsInt(x FX) bool {
	return x == (x &^ mask)
}
