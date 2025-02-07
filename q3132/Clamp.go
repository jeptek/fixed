package q3132

func Clamp(x, min, max FX) FX {
	switch {
	case x == NaN:
		return NaN
	case Le(x, min):
		return min
	case Ge(x, max):
		return max
	default:
		return x
	}
}
