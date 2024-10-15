package q3132

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
//	Leq(-0, +0) = true
//	Leq(+0, -0) = false
//	Leq(+0, 0) = false
//	Leq(0, +0) = true
//	Leq(-0, 0) = true
//	Leq(0, -0) = false
//	Leq(x, y) = x <= y
func Leq(x, y FX) bool {
	if x == NaN || y == NaN {
		panic(ErrNaNComparison)
	} else if x == y {
		if x == Infs || x == -Infs {
			panic(ErrInfsComparison)
		} else if x == Inf || x == -Inf {
			panic(ErrInfComparison)
		}
		return true
	} else if x == Infs {
		return 0 < y
	} else if x == -Infs {
		return 0 <= y
	} else if y == Infs {
		return x <= 0
	} else if y == -Infs {
		return x < 0
	} else {
		return x <= y
	}
}
