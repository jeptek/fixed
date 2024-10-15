package q3132

import "testing"

func TestModf(t *testing.T) {
	type op struct {
		in   FX
		want [2]FX
	}

	cases := []op{
		{in: NaN, want: [2]FX{NaN, NaN}},
		{in: Inf, want: [2]FX{Inf, NaN}},
		{in: -Inf, want: [2]FX{-Inf, NaN}},
		{in: Infs, want: [2]FX{0, Iota}},
		{in: -Infs, want: [2]FX{0, -Iota}},
		{in: -MaxValue, want: [2]FX{-9223372032559808512, -4294967293}},
		{in: MaxValue, want: [2]FX{9223372032559808512, 4294967293}},
		{in: -Pi, want: [2]FX{-3 * One, -608135816}},
		{in: -PiInv, want: [2]FX{0, -PiInv}},
		{in: Pi, want: [2]FX{3 * One, 608135816}},
		{in: PiInv, want: [2]FX{0, PiInv}},
	}

	for _, c := range cases {
		got0, got1 := Modf(c.in)
		if got0 != c.want[0] || got1 != c.want[1] {
			t.Errorf(
				"%s, %s := Modf(%s): want %s, %s",
				got0.String(),
				got1.String(),
				c.in.String(),
				c.want[0].String(),
				c.want[1].String(),
			)
		}
	}

}
