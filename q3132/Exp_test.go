package q3132

import "testing"

func BenchmarkExp_Pi(b *testing.B) {
	Exp(Pi)
}

func BenchmarkExp_PiInv(b *testing.B) {
	Exp(PiInv)
}

func BenchmarkExp_Two(b *testing.B) {
	Exp(Two)
}

func BenchmarkExp_Four(b *testing.B) {
	const four = 4 * One
	Exp(four)
}

func TestExp(t *testing.T) {
	type op struct {
		in   FX
		want FX
	}

	cases := []op{
		{in: NaN, want: NaN},
		{in: Inf, want: Inf},
		{in: -Inf, want: 0},
		{in: Zero, want: One},
		{in: Infs, want: One},
		{in: -Infs, want: One},
		{in: One, want: E},
		{in: -One, want: 1580030169},                 // 0.367879441240802
		{in: -Pi, want: 185602366},                   // 0.04321391826
		{in: -PiInv, want: 3124061927},               //0.72737734929
		{in: Pi, want: 99388518112},                  // 23.1406926328
		{in: PiInv, want: 5904730605},                // 1.37480222744
		{in: 4 * Pi, want: 1231587515564032},         // 286751.313964843750000
		{in: 5 * One, want: 637429665536},            // 148.413159103
		{in: 92288378601, want: Inf},                 // overflow
		{in: -95265423097, want: Infs},               // underflow
		{in: 92288378600, want: 9223371528974893056}, // overflow-1
		{in: -95265423096, want: 178956971},          // underflow+1
	}

	for _, c := range cases {
		if got, want := Exp(c.in), c.want; got != want {
			t.Errorf(
				"Exp(%s): got (%s)[%d] want (%s)[%d]",
				c.in,
				got,
				got,
				want,
				want,
			)
		}
	}
}
