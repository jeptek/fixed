package q3132

// Min returns smallest value in values, will panic if no values are provided
//
// If NaN values are encoutered, this function will return NaN.
//
// Comparisions are done with Le function
func Min(values ...FX) FX {
	if len(values) == 0 {
		panic(ErrZeroLengthArray)
	}
	min := values[0]
	if min == NaN {
		return NaN
	}
	for i := 1; i < len(values); i++ {
		v := values[i]
		if v == NaN {
			return NaN
		} else if v != min && Le(v, min) {
			min = v
		}
	}
	return min
}
