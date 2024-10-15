package q3132

import "testing"

func BenchmarkDiv_Pi_Pi(b *testing.B) {
	Div(Pi, Pi)
}
func BenchmarkDiv_Pi_PiInv(b *testing.B) {
	Div(Pi, PiInv)
}

func BenchmarkDiv_Pi_Two(b *testing.B) {
	Div(Pi, Two)
}

func BenchmarkDiv_PiInv_Pi(b *testing.B) {
	Div(PiInv, Pi)
}
func BenchmarkDiv_PiInv_PiInv(b *testing.B) {
	Div(PiInv, PiInv)
}

func BenchmarkDiv_PiInv_Two(b *testing.B) {
	Div(PiInv, Two)
}

func BenchmarkDiv_Two_Two(b *testing.B) {
	Div(Two, Two)
}

func BenchmarkDiv_Two_Pi(b *testing.B) {
	Div(Two, Pi)
}

func BenchmarkDiv_Two_PiInv(b *testing.B) {
	Div(Two, PiInv)
}

func BenchmarkDiv_NeqTwo_Two(b *testing.B) {
	Div(-Two, Two)
}

func BenchmarkDiv_NeqTwo_Pi(b *testing.B) {
	Div(-Two, Pi)
}

func BenchmarkDiv_NeqTwo_PiInv(b *testing.B) {
	Div(-Two, PiInv)
}

func BenchmarkDiv_NeqTwo_NeqTwo(b *testing.B) {
	Div(-Two, -Two)
}

func BenchmarkDiv_NeqTwo_NeqPi(b *testing.B) {
	Div(-Two, -Pi)
}

func BenchmarkDiv_NeqTwo_NeqPiInv(b *testing.B) {
	Div(-Two, -PiInv)
}

func TestDiv(t *testing.T) {

	type op struct {
		x    FX
		y    FX
		want FX
	}

	cases := []op{
		{x: PiInv, y: PiInv, want: One},
		{x: PiInv, y: -PiInv, want: -One},
		{x: MaxValue, y: Iota, want: Inf},
		{x: MaxValue, y: -Iota, want: -Inf},
		{x: MaxValue, y: One - Iota, want: Inf},
		{x: MaxValue, y: -One + Iota, want: -Inf},
		{x: MaxValue, y: PiInv, want: Inf},
		{x: -MaxValue, y: Iota, want: -Inf},
		{x: -MaxValue, y: -Iota, want: Inf},
		{x: F(0.5), y: MaxValue, want: Iota},
		{x: F(0.5), y: -MaxValue, want: -Iota},
		{x: F(-0.5), y: MaxValue, want: -Iota},
		{x: F(-0.5), y: -MaxValue, want: Iota},
		{x: F(0.25), y: MaxValue, want: Iota},          // round up to iota
		{x: F(0.25), y: -MaxValue, want: -Iota},        // round down to iota
		{x: F(-0.25), y: MaxValue, want: -Iota},        // round down to iota
		{x: F(-0.25), y: -MaxValue, want: Iota},        // round up to iota
		{x: F(0.25) - Iota, y: MaxValue, want: Infs},   // "some number between Iota & zero"
		{x: F(0.25) - Iota, y: -MaxValue, want: -Infs}, // "some number between Iota & zero"
		{x: F(-0.25) + Iota, y: MaxValue, want: -Infs}, // "some number between Iota & zero"
		{x: F(-0.25) + Iota, y: -MaxValue, want: Infs}, // "some number between Iota & zero"
		{x: F(0.015625), y: F(0.000244140625), want: I(64)},
		{x: F(1.44140625), y: F(1.44140625), want: One},
		{x: F(1.515625), y: F(1.531250), want: 4251141099},             //0.989795918
		{x: 5115377064673848749, y: 5115377064673848750, want: ONE},    //0.999999999767169
		{x: One, y: 6203371737317376787, want: 3},                      //0.000000692
		{x: One, y: 6203371737317376, want: 2974},                      //0.000000692
		{x: F(1234.5), y: F(-8888.875), want: -596491359},              //-0.138881467
		{x: F(1234.5), y: F(8888.875), want: 596491359},                //0.138881467
		{x: F(11.17823115), y: F(-5.8754541), want: -8171306660},       //-1.902530588908121
		{x: F(13.33761137), y: F(5.87523451), want: 9750181809},        //2.270141106
		{x: F(0.500244140625), y: F(0.500732421875), want: 4290779127}, //0.999024866
		{x: F(1.25), y: F(1.5), want: 3579139413},                      //0.833333333
		{x: F(1.25), y: I(-4), want: -1342177280},                      //-0.312500000
		{x: F(-1.25), y: I(4), want: -1342177280},                      //-0.312500000
		{x: F(-1.25), y: I(-4), want: 1342177280},                      //0.312500000
		{x: F(1.25), y: I(4), want: 1342177280},                        //0.312500000
		{x: F(1.44140625), y: F(1.441650390625), want: 4294239952},     //0.999830652028322
		{x: Pi, y: E, want: 4963811170},                                //1.155727350
		{x: Pi, y: -E, want: -4963811170},                              //-1.155727350
		{x: E, y: Pi, want: 3716246134},                                //0.865255979
		{x: E, y: -Pi, want: -3716246134},                              //-0.865255979
		{x: E, y: PiInv, want: 36677879205},                            //8.539736484
		{x: E, y: -PiInv, want: -36677879205},                          //-8.539736484
	}
	for i, c := range cases {
		if got, want := Div(c.x, c.y), c.want; got != want {
			t.Errorf("\n\t\t[%d]: Div(%d,%d): got %d want %d  >>> diff: %.15f%%\n\t\t\t%s/%s=%s : got: %s",
				i,
				c.x,
				c.y,
				got,
				want,
				100*float64(got-want)/float64(want),
				c.x,
				c.y,
				want,
				got,
			)
		}
	}
}

func TestDivSpecial(t *testing.T) {
	var test = func(t *testing.T, x, y, want FX) {
		if got, want := Div(x, y), want; got != want {
			t.Errorf("\n\t\tDiv(%d,%d): got %d want %d  >>> diff: %.15f%%\n\t\t\t%s/%s=%s : got: %s",
				x,
				y,
				got,
				want,
				100*float64(got-want)/float64(want),
				x,
				y,
				want,
				got,
			)
		}
	}
	test(t, Pi, NaN, NaN)
	test(t, NaN, Pi, NaN)
	test(t, One, Pi, PiInv)
	test(t, Zero, Zero, NaN)
	test(t, Zero, -Inf, Zero)
	test(t, Zero, Inf, Zero)
	test(t, Zero, Pi, Zero)

	test(t, -Infs, Inf, Zero)
	test(t, -Infs, -Inf, Zero)
	test(t, -Infs, -Infs, NaN)
	test(t, -Infs, Infs, NaN)
	test(t, -Infs, -3*Pi, Infs)
	test(t, -Infs, 3*Pi, -Infs)
	test(t, Infs, Inf, Zero)
	test(t, Infs, Inf, Zero)
	test(t, Infs, -Infs, NaN)
	test(t, Infs, Infs, NaN)
	test(t, Infs, -3*Pi, -Infs)
	test(t, Infs, 3*Pi, Infs)

	test(t, -Inf, -Inf, NaN)
	test(t, -Inf, Inf, NaN)
	test(t, -Inf, -3*Pi, Inf)
	test(t, -Inf, 3*Pi, -Inf)
	test(t, -Inf, Zero, NaN)

	test(t, Inf, -Inf, NaN)
	test(t, Inf, Inf, NaN)
	test(t, Inf, -3*Pi, -Inf)
	test(t, Inf, 3*Pi, Inf)
	test(t, Inf, Zero, NaN)

	test(t, Pi, Infs, Inf)
	test(t, Pi, Zero, NaN)
	test(t, Pi, -Infs, -Inf)
	test(t, Pi, One, Pi)
	test(t, Pi, -Inf, Zero)
	test(t, Pi, Inf, Zero)

	test(t, -Pi, Infs, -Inf)
	test(t, -Pi, Zero, NaN)
	test(t, -Pi, -Infs, Inf)
	test(t, -Pi, One, -Pi)
	test(t, -Pi, -Inf, Zero)
	test(t, -Pi, Inf, Zero)
}
