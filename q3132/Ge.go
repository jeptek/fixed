package q3132

// Ge returns x>y
//
// Special cases are:
//
//	Ge(NaN, y) => panic(ErrNaNComparison)
//	Ge(x, NaN) => panic(ErrNaNComparison)
//	Ge(+0, +0) =>  panic(ErrInfsComparison)
//	Ge(-0, -0) =>  panic(ErrInfsComparison)
//	Ge(+Inf, +Inf) => panic(ErrInfComparison)
//	Ge(-Inf, -Inf) => panic(ErrInfComparison)
//	Ge(-0, +0) = false
//	Ge(+0, -0) = true
//	Ge(+0, 0) = true
//	Ge(0, +0) = false
//	Ge(-0, 0) = false
//	Ge(0, -0) = true
//	Ge(x, y) = x > y
func Ge(x, y FX) bool {
	if x == NaN || y == NaN {
		panic(ErrNaNComparison)
	} else if x == y {
		if x == Infs || x == -Infs {
			panic(ErrInfsComparison)
		} else if x == Inf || x == -Inf {
			panic(ErrInfComparison)
		}
		return false
	} else if x == Infs {
		return 0 >= y
	} else if x == -Infs {
		return 0 > y
	} else if y == Infs {
		return x > 0
	} else if y == -Infs {
		return x >= 0
	} else {
		return x > y
	}
}
