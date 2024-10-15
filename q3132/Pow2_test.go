package q3132

import "testing"

func BenchmarkPow2_Pi(b *testing.B) {
	Pow2(Pi)
}

func BenchmarkPow2_NeqPi(b *testing.B) {
	Pow2(-Pi)
}

func BenchmarkPow2_PiInv(b *testing.B) {
	Pow2(PiInv)
}

func BenchmarkPow2_NeqPiInv(b *testing.B) {
	Pow2(-PiInv)
}

func BenchmarkPow2_Two(b *testing.B) {
	Pow2(Two)
}

func BenchmarkPow2_NeqTwo(b *testing.B) {
	Pow2(-Two)
}

func BenchmarkPow2_Four(b *testing.B) {
	const four = 4 * One
	Pow2(four)
}

func BenchmarkPow2_NeqFour(b *testing.B) {
	const four = -4 * One
	Pow2(four)
}

func TestPow2(t *testing.T) {
	type op struct {
		in   FX
		want FX
	}

	cases := []op{
		{in: NaN, want: NaN},
		{in: Zero, want: One},
		{in: -Infs, want: One},
		{in: Infs, want: One},
		{in: -Inf, want: Zero},
		{in: Inf, want: Inf},
		{in: One, want: Two},
		{in: 5 * One, want: I(32)},
		{in: 0x0000001F00000000, want: Inf},   //overflow
		{in: -0x0000002000000000, want: Infs}, //underflow
		{in: -Pi, want: 486683070},            //0.11331473229676088
		{in: Pi, want: 37902991144},           //8.824977827076287
		{in: PiInv, want: 5355261530},         //1.246868988964707
		{in: -PiInv, want: 3444601906},        //0.802008878905326
		{in: Iota, want: 4294967297},          //1.000000000232831
		{in: -Iota, want: 4294967295},         //0.999999999767169
	}

	for i, c := range cases {
		if got, want := Pow2(c.in), c.want; got != want {
			t.Errorf("\n\t\t[%d]: Pow2(%d): got %d want %d  >>> diff: %.15f%%\n\t\t\t2**%s=%s : got: %s",
				i,
				c.in,
				got,
				want,
				100*float64(got-want)/float64(want),
				c.in,
				want,
				got,
			)
		}
	}
}
