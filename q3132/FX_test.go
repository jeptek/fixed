package q3132

import (
	"math"
	"testing"
)

func TestFX_String(t *testing.T) {
	type op struct {
		in   FX
		want string
	}

	cases := []op{
		{in: -163245643 * One, want: "-163245643"},
		{in: -95265423096, want: "-22.180709777399898"},
		{in: -95265423097, want: "-22.180709777632728"},
		{in: -Iota, want: "-0.000000000232831"},
		{in: -Inf, want: "-Inf"},
		{in: -Infs, want: "-0"},
		{in: -MaxValue, want: "-2147483647.999999999301508"},
		{in: -One, want: "-1"},
		{in: -Pi, want: "-3.141592653468251"},
		{in: -PiInv, want: "-0.318309886148199"},
		{in: 17845 * One, want: "17845"},
		{in: 4 * Pi, want: "12.566370613873005"},
		{in: 5 * One, want: "5"},
		{in: 92288378600, want: "21.487562591210008"},
		{in: 92288378601, want: "21.487562591442838"},
		{in: Iota, want: "0.000000000232831"},
		{in: Inf, want: "+Inf"},
		{in: Infs, want: "+0"},
		{in: MaxValue, want: "2147483647.999999999301508"},
		{in: NaN, want: "NaN"},
		{in: One, want: "1"},
		{in: Pi, want: "3.141592653468251"},
		{in: PiInv, want: "0.318309886148199"},
		{in: Zero, want: "0"},
	}

	for _, c := range cases {
		if got, want := c.in.String(), c.want; got != want {
			t.Errorf(
				"String(%d): got [%s] want [%s]",
				c.in,
				got,
				want,
			)
		}
	}
}

func TestFX_Value(t *testing.T) {
	const float64EqualityThreshold = 1e-10

	var almostEqual = func(a, b float64) bool {
		return math.Abs(a-b) <= float64EqualityThreshold
	}

	type op struct {
		in   FX
		want float64
	}

	cases := []op{
		{in: NaN, want: math.NaN()},
		{in: -163245643 * One, want: -163245643},
		{in: -95265423096, want: -22.180709777399898},
		{in: -95265423097, want: -22.18070977763272},
		{in: -Iota, want: -0.000000000232831},
		{in: -Inf, want: math.Inf(-1)},
		{in: -Infs, want: -math.SmallestNonzeroFloat64},
		{in: -MaxValue, want: -2147483648},
		{in: -One, want: -1},
		{in: -Pi, want: -3.141592653468251},
		{in: -PiInv, want: -0.318309886148199},
		{in: 17845 * One, want: 17845},
		{in: 4 * Pi, want: 12.566370613873005},
		{in: 5 * One, want: 5},
		{in: 92288378600, want: 21.487562591210008},
		{in: 92288378601, want: 21.487562591442838},
		{in: Iota, want: 0.000000000232831},
		{in: Inf, want: math.Inf(1)},
		{in: Infs, want: math.SmallestNonzeroFloat64},
		{in: MaxValue, want: 2147483648},
		{in: One, want: 1},
		{in: Pi, want: 3.141592653468251},
		{in: PiInv, want: 0.318309886148199},
		{in: Zero, want: 0},
	}

	for _, c := range cases {
		if got, want := c.in.Value(), c.want; !almostEqual(got, want) {
			if math.IsNaN(got) && math.IsNaN(want) {
				// Cannot comparte NaN directly
			} else if math.IsInf(got, 1) && math.IsInf(want, 1) {
				// Cannot comparte Inf directly
			} else if math.IsInf(got, -1) && math.IsInf(want, -1) {
				// Cannot comparte Inf directly
			} else {
				t.Errorf(
					"Value(%s): got [%.15f] want [%.15f]",
					c.in,
					got,
					want,
				)
			}
		}
	}
}

func TestFX_Floor(t *testing.T) {
	type op struct {
		in   FX
		want int
	}

	cases := []op{
		{in: -163245643 * One, want: -163245643},
		{in: -95265423096, want: -23},
		{in: -95265423097, want: -23},
		{in: -Iota, want: -1},
		{in: -Inf, want: -2147483648},
		{in: -Infs, want: -1},
		{in: -MaxValue, want: -2147483648},
		{in: -One, want: -1},
		{in: -Pi, want: -4},
		{in: -PiInv, want: -1},
		{in: 17845 * One, want: 17845},
		{in: 4 * Pi, want: 12},
		{in: 5 * One, want: 5},
		{in: 92288378600, want: 21},
		{in: 92288378601, want: 21},
		{in: Iota, want: 0},
		{in: Inf, want: 2147483647},
		{in: Infs, want: 0},
		{in: MaxValue, want: 2147483647},
		{in: One, want: 1},
		{in: Pi, want: 3},
		{in: PiInv, want: 0},
		{in: Zero, want: 0},
	}

	for _, c := range cases {
		if got, want := c.in.Floor(), c.want; got != want {
			t.Errorf(
				"Floor(%s): got [%d] want [%d]",
				c.in,
				got,
				want,
			)
		}
	}
}

func TestFX_Floor_NaN(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	NaN.Floor()
}

func TestFX_Round(t *testing.T) {
	type op struct {
		in   FX
		want int
	}

	cases := []op{
		{in: -163245643 * One, want: -163245643},
		{in: -95265423096, want: -22},
		{in: -95265423097, want: -22},
		{in: -Iota, want: 0},
		{in: -Inf, want: -2147483648},
		{in: -Infs, want: 0},
		{in: -MaxValue, want: -2147483648},
		{in: -One, want: -1},
		{in: -Pi, want: -3},
		{in: -PiInv, want: 0},
		{in: 17845 * One, want: 17845},
		{in: 4 * Pi, want: 13},
		{in: 5 * One, want: 5},
		{in: 92288378600, want: 21},
		{in: 92288378601, want: 21},
		{in: Iota, want: 0},
		{in: Inf, want: 2147483647},
		{in: Infs, want: 0},
		{in: MaxValue, want: 2147483647},
		{in: MaxValue - Iota, want: 2147483647},
		{in: One, want: 1},
		{in: Pi, want: 3},
		{in: PiInv, want: 0},
		{in: Zero, want: 0},
	}

	for _, c := range cases {
		if got, want := c.in.Round(), c.want; got != want {
			t.Errorf(
				"Round(%s): got [%d] want [%d]",
				c.in,
				got,
				want,
			)
		}
	}
}

func TestFX_Round_NaN(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	NaN.Round()
}

func TestFX_Ceil(t *testing.T) {
	type op struct {
		in   FX
		want int
	}

	cases := []op{
		{in: -163245643 * One, want: -163245643},
		{in: -95265423096, want: -22},
		{in: -95265423097, want: -22},
		{in: -Iota, want: 0},
		{in: -Inf, want: -2147483648},
		{in: -Infs, want: 0},
		{in: -MaxValue, want: -2147483648},
		{in: -One, want: -1},
		{in: -Pi, want: -3},
		{in: -PiInv, want: 0},
		{in: 17845 * One, want: 17845},
		{in: 4 * Pi, want: 13},
		{in: 5 * One, want: 5},
		{in: 92288378600, want: 22},
		{in: 92288378601, want: 22},
		{in: Iota, want: 1},
		{in: Inf, want: 2147483647},
		{in: Infs, want: 1},
		{in: MaxValue, want: 2147483647},
		{in: MaxValue - Iota, want: 2147483647},
		{in: One, want: 1},
		{in: Pi, want: 4},
		{in: PiInv, want: 1},
		{in: Zero, want: 0},
	}

	for _, c := range cases {
		if got, want := c.in.Ceil(), c.want; got != want {
			t.Errorf(
				"Ceil(%s): got [%d] want [%d]",
				c.in,
				got,
				want,
			)
		}
	}
}

func TestFX_Ceil_NaN(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	NaN.Ceil()
}

func TestFX_FillCoverageForTestedMethods(t *testing.T) {
	// just to fill coverage (functions already tested directly):
	x := Pi
	y := PiInv
	if got, want := x.Add(y), Add(x, y); got != want {
		t.Errorf("Add(%s,%s): got [%d](%s) want [%d](%s)",
			x,
			y,
			got,
			got,
			want,
			want,
		)
	}

	if got, want := x.Sub(y), Sub(x, y); got != want {
		t.Errorf("Sub(%s,%s): got [%d](%s) want [%d](%s)",
			x,
			y,
			got,
			got,
			want,
			want,
		)
	}

	if got, want := x.Mul(y), Mul(x, y); got != want {
		t.Errorf("Mul(%s,%s): got [%d](%s) want [%d](%s)",
			x,
			y,
			got,
			got,
			want,
			want,
		)
	}

	if got, want := x.Div(y), Div(x, y); got != want {
		t.Errorf("Div(%s,%s): got [%d](%s) want [%d](%s)",
			x,
			y,
			got,
			got,
			want,
			want,
		)
	}

	if got, want := x.Le(y), Le(x, y); got != want {
		t.Errorf("Le(%s,%s): got %t want %t",
			x,
			y,
			got,
			want,
		)
	}

	if got, want := x.Leq(y), Leq(x, y); got != want {
		t.Errorf("Leq(%s,%s): got %t want %t",
			x,
			y,
			got,
			want,
		)
	}

	if got, want := x.Ge(y), Ge(x, y); got != want {
		t.Errorf("Ge(%s,%s): got %t want %t",
			x,
			y,
			got,
			want,
		)
	}

	if got, want := x.Geq(y), Geq(x, y); got != want {
		t.Errorf("Geq(%s,%s): got %t want %t",
			x,
			y,
			got,
			want,
		)
	}

	if got, want := x.Sqrt(), Sqrt(x); got != want {
		t.Errorf("Sqrt(%s): got [%d](%s) want [%d](%s)",
			x,
			got,
			got,
			want,
			want,
		)
	}

	if got, want := x.Exp(), Exp(x); got != want {
		t.Errorf("Exp(%s): got [%d](%s) want [%d](%s)",
			x,
			got,
			got,
			want,
			want,
		)
	}

	if got, want := x.Ln(), Ln(x); got != want {
		t.Errorf("Ln(%s): got [%d](%s) want [%d](%s)",
			x,
			got,
			got,
			want,
			want,
		)
	}

	if got, want := x.Log2(), Log2(x); got != want {
		t.Errorf("Log2(%s): got [%d](%s) want [%d](%s)",
			x,
			got,
			got,
			want,
			want,
		)
	}

	if got, want := x.Pow(y), Pow(x, y); got != want {
		t.Errorf("Pow(%s): got [%d](%s) want [%d](%s)",
			x,
			got,
			got,
			want,
			want,
		)
	}

}
