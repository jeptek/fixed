package q3132

import "testing"

func TestIsReal(t *testing.T) {
	var test = func(t *testing.T, x FX, want bool) {
		if got, want := IsReal(x), want; got != want {
			t.Errorf("\n\t\tIsReal(%s): got %t want %t",
				x,
				got,
				want,
			)
		}
	}
	test(t, NaN, false)
	test(t, Inf, false)
	test(t, -Inf, false)
	test(t, Infs, false)
	test(t, -Infs, false)
	test(t, Pi, true)
	test(t, One, true)
	test(t, F(-323.232), true)
}
