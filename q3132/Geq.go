package q3132

// Geq returns x>=y
//
// Special cases are:
//
//	Geq(NaN, y) => panic(ErrNaNComparison)
//	Geq(x, NaN) => panic(ErrNaNComparison)
//	Geq(+0, +0) => panic(ErrInfsComparison)
//	Geq(-0, -0) => panic(ErrInfsComparison)
//	Geq(+Inf, +Inf) => panic(ErrInfComparison)
//	Geq(-Inf, -Inf) => panic(ErrInfComparison)
//	Geq(-0, +0) = false
//	Geq(+0, -0) = true
//	Geq(+0, 0) = true
//	Geq(0, +0) = false
//	Geq(-0, 0) = false
//	Geq(0, -0) = true
//	Geq(x, y) = x >= y
func Geq(x, y FX) bool {
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
		return 0 >= y
	} else if x == -Infs {
		return 0 > y
	} else if y == Infs {
		return x > 0
	} else if y == -Infs {
		return x >= 0
	} else {
		return x >= y
	}
}
