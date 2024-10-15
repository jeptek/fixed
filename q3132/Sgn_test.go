package q3132

import "testing"

func TestSgn(t *testing.T) {
	var test = func(t *testing.T, x, want FX) {
		if got, want := Sgn(x), want; got != want {
			t.Errorf("\n\t\tSgn(%s): got (%s)[%d] want (%s)[%d]",
				x,
				got,
				got,
				want,
				want,
			)
		}
	}
	test(t, NaN, NaN)
	test(t, One, One)
	test(t, Zero, Zero)
	test(t, -Infs, -One)
	test(t, Infs, One)
	test(t, -Inf, -One)
	test(t, Inf, One)

	test(t, Pi, One)
	test(t, -Pi, -One)
}

func Test_sgn(t *testing.T) {
	var test = func(t *testing.T, x, want FX) {
		if got, want := sgn(x), want; got != want {
			t.Errorf("\n\t\tsgn(%s): got (%s)[%d] want (%s)[%d]",
				x,
				got,
				got,
				want,
				want,
			)
		}
	}
	test(t, NaN, -1) // expected for internal calls
	test(t, One, 1)
	test(t, Zero, 0)
	test(t, -Infs, -1)
	test(t, Infs, 1)
	test(t, -Inf, -1)
	test(t, Inf, 1)
	test(t, Pi, 1)
	test(t, -Pi, -1)
}
