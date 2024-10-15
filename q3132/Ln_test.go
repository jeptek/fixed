package q3132

import "testing"

func BenchmarkLn_Pi(b *testing.B) {
	Ln(Pi)
}

func BenchmarkLn_PiInv(b *testing.B) {
	Ln(PiInv)
}

func BenchmarkLn_Two(b *testing.B) {
	Ln(Two)
}

func BenchmarkLn_Four(b *testing.B) {
	const four = 4 * One
	Ln(four)
}

func TestLn(t *testing.T) {
	type op struct {
		x    FX
		want FX
	}

	cases := []op{
		{x: MaxValue, want: 92288378625}, // 21.487562597030774
		{x: Zero, want: -Inf},
		{x: One, want: Zero},
		{x: -Pi, want: NaN},
		{x: Inf, want: Inf},
		{x: -Inf, want: NaN},
		{x: Infs, want: -Inf},
		{x: -Infs, want: NaN},
		{x: 50 * Pi, want: 21718588292}, //5.056752891279757
		{x: Pi, want: 4916577427},       //1.144729887
		{x: PiInv, want: -4916577426},   //-1.144729886669666
		{x: 5 * One, want: 6912483197},  //1.609437912004068
	}

	for _, c := range cases {
		if got, want := Ln(c.x), c.want; got != want {
			t.Errorf(
				"Ln(%s): got (%s)[%d] want (%s)[%d]",
				c.x,
				got,
				got,
				want,
				want,
			)
		}
	}
}
