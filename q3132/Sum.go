package q3132

// Sum uses Add repeatedly to combine values.
//
// Special cases:
//
//	Sum(len(values) == 0) = 0
func Sum(values ...FX) FX {
	if len(values) == 0 {
		return Zero
	}
	res := values[0]
	for i := 1; i < len(values); i++ {
		res = Add(res, values[i])
	}
	return res
}
