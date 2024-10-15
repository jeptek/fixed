package q3132

import "testing"

func TestLeq(t *testing.T) {
	var test = func(t *testing.T, x, y FX, want bool) {
		if got, want := Leq(x, y), want; got != want {
			t.Errorf("Leq(%s,%s): got %t want %t",
				x,
				y,
				got,
				want,
			)
		}
	}
	test(t, One, Pi, true)
	test(t, One, One, true)
	test(t, -One, Inf, true)
	test(t, -One, -Inf, false)

	test(t, Zero, Zero, true)
	test(t, Zero, -Inf, false)
	test(t, Zero, Inf, true)
	test(t, Zero, -Infs, false)
	test(t, Zero, Infs, true)
	test(t, Zero, Pi, true)
	test(t, Zero, Iota, true)
	test(t, Zero, -Iota, false)

	test(t, -Infs, Inf, true)
	test(t, -Infs, -Inf, false)
	//test(t, -Infs, -Infs, NaN)
	test(t, -Infs, Infs, true)
	test(t, -Infs, -3*Pi, false)
	test(t, -Infs, 3*Pi, true)
	test(t, -Infs, Iota, true)
	test(t, -Infs, -Iota, false)
	test(t, -Infs, Zero, true)
	test(t, -Infs, I(18), true)
	test(t, -Infs, I(7531), true)

	test(t, Infs, Inf, true)
	test(t, Infs, -Inf, false)
	test(t, Infs, -Infs, false)
	//test(t, Infs, Infs, One)
	test(t, Infs, -3*Pi, false)
	test(t, Infs, 3*Pi, true)
	test(t, Infs, Iota, true)
	test(t, Infs, -Iota, false)

	test(t, -Iota, -Iota, true)
	test(t, -Iota, Inf, true)
	test(t, -Iota, -Inf, false)
	test(t, -Iota, -Infs, true)
	test(t, -Iota, Infs, true)
	test(t, -Iota, -3*Pi, false)
	test(t, -Iota, 3*Pi, true)

	test(t, Iota, Iota, true)
	test(t, Iota, Inf, true)
	test(t, Iota, -Inf, false)
	test(t, Iota, -Infs, false)
	test(t, Iota, Infs, false)
	test(t, Iota, -3*Pi, false)
	test(t, Iota, 3*Pi, true)

	test(t, -Inf, Inf, true)
	test(t, -Inf, Infs, true)
	test(t, -Inf, -Infs, true)
	test(t, -Inf, -3*Pi, true)
	test(t, -Inf, 3*Pi, true)
	test(t, -Inf, Zero, true)

	test(t, Inf, -Inf, false)
	test(t, Inf, Infs, false)
	test(t, Inf, -Infs, false)
	test(t, Inf, -3*Pi, false)
	test(t, Inf, 3*Pi, false)
	test(t, Inf, Zero, false)

	test(t, Pi, Pi, true)
	test(t, Pi, Infs, false)
	test(t, Pi, Zero, false)
	test(t, Pi, -Infs, false)
	test(t, Pi, One, false)
	test(t, Pi, 2*Pi, true)
	test(t, Pi, PiInv, false)
	test(t, Pi, -Inf, false)
	test(t, Pi, Inf, true)

	test(t, -Pi, -Pi, true)
	test(t, -Pi, Infs, true)
	test(t, -Pi, Zero, true)
	test(t, -Pi, -Infs, true)
	test(t, -Pi, One, true)
	test(t, -Pi, -One, true)
	test(t, -Pi, -Inf, false)
	test(t, -Pi, Inf, true)

	test(t, PiInv, PiInv, true)
	test(t, PiInv, Infs, false)
	test(t, PiInv, Zero, false)
	test(t, PiInv, -Infs, false)
	test(t, PiInv, One, true)
	test(t, PiInv, -Inf, false)
	test(t, PiInv, Inf, true)

	test(t, -PiInv, -PiInv, true)
	test(t, -PiInv, Infs, true)
	test(t, -PiInv, Zero, true)
	test(t, -PiInv, -Infs, true)
	test(t, -PiInv, One, true)
	test(t, -PiInv, -Inf, false)
	test(t, -PiInv, Inf, true)
}

func TestLeq_Inf_Inf(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	Leq(Inf, Inf)
}

func TestLeq_NegInf_NegInf(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	Leq(-Inf, -Inf)
}

func TestLeq_Infs_Infs(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	Leq(Infs, Infs)
}

func TestLeq_NegInfs_NegInfs(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	Leq(-Infs, -Infs)
}

func TestLeq_NaN(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	Leq(NaN, -Infs)
}
