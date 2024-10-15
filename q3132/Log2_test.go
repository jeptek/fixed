package q3132

import "testing"

func BenchmarkLog2_Pi(b *testing.B) {
	Log2(Pi)
}

func BenchmarkLog2_PiInv(b *testing.B) {
	Log2(PiInv)
}

func BenchmarkLog2_Two(b *testing.B) {
	Log2(Two)
}

func BenchmarkLog2_Four(b *testing.B) {
	const four = 4 * One
	Log2(four)
}

func TestLog2(t *testing.T) {
	type op struct {
		in   FX
		want FX
	}

	cases := []op{
		{in: NaN, want: NaN},
		{in: Inf, want: Inf},
		{in: -Inf, want: NaN},
		{in: Zero, want: -Inf},
		{in: Infs, want: -Inf},
		{in: -Infs, want: NaN},
		{in: One, want: Zero},
		{in: -One, want: NaN},
		{in: -Pi, want: NaN},
		{in: -PiInv, want: NaN},
		{in: I(32), want: I(5)},
		{in: I(17), want: 17555519226},     //4.087462841067463
		{in: Pi, want: 7093121865},         // 1.651496129343286
		{in: PiInv, want: -7093121866},     // -1.651496129576117
		{in: 4 * Pi, want: 15683056457},    // 3.651496129343286
		{in: 5 * One, want: 9972605231},    // 2.321928094839677
		{in: MaxValue, want: 133143986176}, // 31
	}

	for _, c := range cases {
		if got, want := Log2(c.in), c.want; got != want {
			t.Errorf(
				"Log2(%s): got (%s)[%d] want (%s)[%d]",
				c.in,
				got,
				got,
				want,
				want,
			)
		}
	}
}
