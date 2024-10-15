package q3132

// Max returns largest value in values, will panic if no values are provided
//
// If NaN values are encoutered, this function will return NaN.
//
// Comparisions are done with Ge function
func Max(values ...FX) FX {
	if len(values) == 0 {
		panic(ErrZeroLengthArray)
	}
	max := values[0]
	if max == NaN {
		return NaN
	}
	for i := 1; i < len(values); i++ {
		v := values[i]
		if v == NaN {
			return NaN
		} else if v != max && Ge(v, max) {
			max = v
		}
	}
	return max
}
