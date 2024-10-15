package q3132

import "testing"

func TestAbs(t *testing.T) {
	var test = func(t *testing.T, x, want FX) {
		if got, want := Abs(x), want; got != want {
			t.Errorf("\n\t\tAbs(%s): got (%s)[%d] want (%s)[%d]",
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
	test(t, -Infs, Infs)
	test(t, Infs, Infs)
	test(t, -Inf, Inf)
	test(t, Inf, Inf)

	test(t, Pi, Pi)
	test(t, -Pi, Pi)
}

func Test_abs(t *testing.T) {
	var test = func(t *testing.T, x, want FX) {
		if got, want := abs(x), want; got != want {
			t.Errorf("\n\t\tabs(%s): got (%s)[%d] want (%s)[%d]",
				x,
				got,
				got,
				want,
				want,
			)
		}
	}
	test(t, NaN, NaN) // expected for internal calls
	test(t, One, One)
	test(t, Zero, Zero)
	test(t, -Infs, Infs)
	test(t, Infs, Infs)
	test(t, -Inf, Inf)
	test(t, Inf, Inf)
	test(t, Pi, Pi)
	test(t, -Pi, Pi)
}
