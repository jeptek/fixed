package q3132

import "testing"

func TestGe(t *testing.T) {
	var test = func(t *testing.T, x, y FX, want bool) {
		if got, want := Ge(x, y), want; got != want {
			t.Errorf("Ge(%s,%s): got %t want %t",
				x,
				y,
				got,
				want,
			)
		}
	}
	test(t, One, Pi, false)
	test(t, -One, Inf, false)
	test(t, -One, -Inf, true)

	test(t, Zero, Zero, false)
	test(t, Zero, -Inf, true)
	test(t, Zero, Inf, false)
	test(t, Zero, -Infs, true)
	test(t, Zero, Infs, false)
	test(t, Zero, Pi, false)
	test(t, Zero, Iota, false)
	test(t, Zero, -Iota, true)

	test(t, -Infs, Inf, false)
	test(t, -Infs, -Inf, true)
	//test(t, -Infs, -Infs, NaN)
	test(t, -Infs, Infs, false)
	test(t, -Infs, -3*Pi, true)
	test(t, -Infs, 3*Pi, false)
	test(t, -Infs, Iota, false)
	test(t, -Infs, -Iota, true)
	test(t, -Infs, Zero, false)
	test(t, -Infs, I(18), false)
	test(t, -Infs, I(7531), false)

	test(t, Infs, Inf, false)
	test(t, Infs, -Inf, true)
	test(t, Infs, -Infs, true)
	//test(t, Infs, Infs, One)
	test(t, Infs, -3*Pi, true)
	test(t, Infs, 3*Pi, false)
	test(t, Infs, Iota, false)
	test(t, Infs, -Iota, true)

	test(t, -Iota, Inf, false)
	test(t, -Iota, -Inf, true)
	test(t, -Iota, -Infs, false)
	test(t, -Iota, Infs, false)
	test(t, -Iota, -3*Pi, true)
	test(t, -Iota, 3*Pi, false)

	test(t, Iota, Inf, false)
	test(t, Iota, -Inf, true)
	test(t, Iota, -Infs, true)
	test(t, Iota, Infs, true)
	test(t, Iota, -3*Pi, true)
	test(t, Iota, 3*Pi, false)

	test(t, -Inf, Inf, false)
	test(t, -Inf, Infs, false)
	test(t, -Inf, -Infs, false)
	test(t, -Inf, -3*Pi, false)
	test(t, -Inf, 3*Pi, false)
	test(t, -Inf, Zero, false)

	test(t, Inf, -Inf, true)
	test(t, Inf, Infs, true)
	test(t, Inf, -Infs, true)
	test(t, Inf, -3*Pi, true)
	test(t, Inf, 3*Pi, true)
	test(t, Inf, Zero, true)

	test(t, Pi, Infs, true)
	test(t, Pi, Zero, true)
	test(t, Pi, -Infs, true)
	test(t, Pi, One, true)
	test(t, Pi, 2*Pi, false)
	test(t, Pi, PiInv, true)
	test(t, Pi, -Inf, true)
	test(t, Pi, Inf, false)

	test(t, -Pi, Infs, false)
	test(t, -Pi, Zero, false)
	test(t, -Pi, -Infs, false)
	test(t, -Pi, One, false)
	test(t, -Pi, -One, false)
	test(t, -Pi, -Inf, true)
	test(t, -Pi, Inf, false)

	test(t, PiInv, Infs, true)
	test(t, PiInv, Zero, true)
	test(t, PiInv, -Infs, true)
	test(t, PiInv, One, false)
	test(t, PiInv, -Inf, true)
	test(t, PiInv, Inf, false)

	test(t, -PiInv, Infs, false)
	test(t, -PiInv, Zero, false)
	test(t, -PiInv, -Infs, false)
	test(t, -PiInv, One, false)
	test(t, -PiInv, -Inf, true)
	test(t, -PiInv, Inf, false)
}

func TestGe_Inf_Inf(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	Ge(Inf, Inf)
}

func TestGe_NegInf_NegInf(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	Ge(-Inf, -Inf)
}

func TestGe_Infs_Infs(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	Ge(Infs, Infs)
}

func TestGe_NegInfs_NegInfs(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	Ge(-Infs, -Infs)
}

func TestGe_NaN(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	Ge(NaN, -Infs)
}
