package q3132

import (
	"testing"
)

func TestMin(t *testing.T) {
	type op struct {
		in   []FX
		want FX
	}

	cases := []op{
		{in: []FX{NaN, MaxValue, Zero, Pi, Inf}, want: NaN},
		{in: []FX{Zero, Pi, Inf, NaN}, want: NaN},
		{in: []FX{Zero, Pi, -Inf, NaN}, want: NaN},
		{in: []FX{Zero, Pi, -Inf, Infs}, want: -Inf},
		{in: []FX{Zero, PiInv, Inf, -Infs}, want: -Infs},
		{in: []FX{Zero, -PiInv, Inf, -Infs}, want: -PiInv},
		{in: []FX{-One, -PiInv, Inf, -Infs}, want: -One},
		{in: []FX{-One, -Infs, -PiInv, -Infs}, want: -One},
		{in: []FX{-One, -Infs, -PiInv, Infs}, want: -One},
		{in: []FX{-One, -Pi, Infs, -PiInv}, want: -Pi},
	}
	for _, c := range cases {
		if got, want := Min(c.in...), c.want; got != want {
			t.Errorf(
				"Min(%s): got (%s)[%d] want (%s)[%d]",
				c.in,
				got,
				got,
				want,
				want,
			)
		}
	}
}

func TestMin_NoArguments(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	Min()
}
