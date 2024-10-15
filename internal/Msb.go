package internal

// Msb returns the minimum number of bits required to represent x; the result is 0 for x == 0.
func Msb(x uint64) int {
	n := 0
	if x >= 1<<32 {
		x >>= 32
		n += 32
	}
	if x >= 1<<16 {
		x >>= 16
		n += 16
	}
	if x >= 1<<8 {
		x >>= 8
		n += 8
	}
	if x >= 1<<4 {
		x >>= 4
		n += 4
	}
	if x >= 1<<2 {
		x >>= 2
		n += 2
	}
	if x >= 1<<1 {
		x >>= 1
		n += 1
	}
	if x >= 1 {
		n += 1
	}
	return n
}
