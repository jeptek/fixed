package q3132

// Constraint for types that are convertible to int64 without flipping sign.
type Int64 interface {
	~uint8 | ~uint16 | ~uint32 | ~int | ~int8 | ~int16 | ~int32 | ~int64
}
