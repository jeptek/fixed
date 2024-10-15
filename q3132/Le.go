package q3132

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
func Le(x, y FX) bool {
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
		return 0 < y
	} else if x == -Infs {
		return 0 <= y
	} else if y == Infs {
		return x <= 0
	} else if y == -Infs {
		return x < 0
	} else {
		return x < y
	}
}
