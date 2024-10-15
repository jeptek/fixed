package q3132

import (
	"testing"
)

func TestMax(t *testing.T) {
	type op struct {
		in   []FX
		want FX
	}

	cases := []op{
		{in: []FX{NaN, MaxValue, Zero, Pi, Inf}, want: NaN},
		{in: []FX{Zero, Pi, Inf, NaN}, want: NaN},
		{in: []FX{Zero, Pi, -Inf, NaN}, want: NaN},
		{in: []FX{Zero, Pi, -Inf, Infs}, want: Pi},
		{in: []FX{Zero, PiInv, -Inf, -Infs}, want: PiInv},
		{in: []FX{Zero, -PiInv, -Inf, -Infs}, want: Zero},
		{in: []FX{-One, -PiInv, -Inf, -Infs}, want: -Infs},
		{in: []FX{-One, -Infs, -PiInv, -Infs}, want: -Infs},
		{in: []FX{-One, -Infs, -PiInv, Infs}, want: Infs},
		{in: []FX{-One, -Pi, Infs, -PiInv}, want: Infs},
	}
	for _, c := range cases {
		if got, want := Max(c.in...), c.want; got != want {
			t.Errorf(
				"Max(%s): got (%s)[%d] want (%s)[%d]",
				c.in,
				got,
				got,
				want,
				want,
			)
		}
	}
}

func TestMax_NoArguments(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	Max()
}
