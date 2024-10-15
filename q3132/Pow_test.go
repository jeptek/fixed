package q3132

import (
	"testing"
)

func BenchmarkPow_Pi_Pi(b *testing.B) {
	Pow(Pi, Pi)
}
func BenchmarkPow_Pi_PiInv(b *testing.B) {
	Pow(Pi, PiInv)
}

func BenchmarkPow_Pi_Two(b *testing.B) {
	Pow(Pi, Two)
}

func BenchmarkPow_PiInv_Pi(b *testing.B) {
	Pow(PiInv, Pi)
}
func BenchmarkPow_PiInv_PiInv(b *testing.B) {
	Pow(PiInv, PiInv)
}

func BenchmarkPow_PiInv_Two(b *testing.B) {
	Pow(PiInv, Two)
}

func BenchmarkPow_Two_Two(b *testing.B) {
	Pow(Two, Two)
}

func BenchmarkPow_Two_Pi(b *testing.B) {
	Pow(Two, Pi)
}

func BenchmarkPow_Two_PiInv(b *testing.B) {
	Pow(Two, PiInv)
}

func BenchmarkPow_NeqTwo_Two(b *testing.B) {
	Pow(-Two, Two)
}

func BenchmarkPow_NeqTwo_Pi(b *testing.B) {
	Pow(-Two, Pi)
}

func BenchmarkPow_NeqTwo_NeqTwo(b *testing.B) {
	Pow(-Two, -Two)
}

func BenchmarkPow_NeqTwo_100(b *testing.B) {
	Pow(-Two, 100)
}

func BenchmarkPow_NeqTwo_101(b *testing.B) {
	Pow(-Two, 101)
}

func BenchmarkPow_NeqTwo_Neq100(b *testing.B) {
	Pow(-Two, -100)
}

func BenchmarkPow_NeqTwo_Neq101(b *testing.B) {
	Pow(-Two, -101)
}

func TestPow(t *testing.T) {
	type op struct {
		x    FX
		y    FX
		want FX
	}

	cases := []op{
		{x: -PiInv, y: PiInv, want: NaN},
		{x: -PiInv, y: -PiInv, want: NaN},
		{x: F(1234.5), y: F(-8888.875), want: Infs},
		{x: F(1234.5), y: F(8888.875), want: Inf},
		{x: F(-14423.35375565), y: I(-17), want: -Infs},
		{x: F(-14423.35375565), y: I(-14), want: Infs},
		{x: 4423816314, y: I(807 - 80), want: Inf},
		{x: F(-1.35375565), y: I(-3), want: -1731167200},               //-0.403068773448467
		{x: F(-1.35375565), y: I(5), want: -19528157612},               //-4.546753506176174
		{x: F(-1.35375565), y: I(12), want: 162720865728},              //37.886403903365135
		{x: F(11.17823115), y: F(-5.8754541), want: 2974},              //0.000000692
		{x: F(13.33761137), y: F(5.87523451), want: 17500775697612800}, //4074716.870117188
		{x: F(0.015625), y: F(0.000244140625), want: 4290608604},       //0.998985163
		{x: F(0.500244140625), y: F(0.500732421875), want: 3036201153}, //0.706920669
		{x: F(1.25), y: F(1.5), want: 6002399271},                      //1.397542486
		{x: F(1.25), y: I(-4), want: 1759218605},                       //0.409600000
		{x: F(1.25), y: I(4), want: 10485759998},                       //2.441406250
		{x: F(1.44140625), y: F(1.44140625), want: 7275043119},         //1.693852972
		{x: F(1.44140625), y: F(1.441650390625), want: 7275692539},     //1.694004177
		{x: F(1.515625), y: F(1.531250), want: 8118784127},             //1.890301734
		{x: Pi, y: E, want: 96461347856},                               //22.459157709
		{x: Pi, y: -E, want: 191234567},                                //0.044525267
		{x: E, y: Pi, want: 99388518048},                               //23.140692629
		{x: E, y: -Pi, want: 185602366},                                //0.043213918
		{x: E, y: PiInv, want: 5904730605},                             //1.374802111
		{x: E, y: -PiInv, want: 3124061927},                            //0.727377411
		{x: PiInv, y: PiInv, want: 2983404509},                         //0.694628001
		{x: PiInv, y: -PiInv, want: 6183118655},                        //1.439619478

	}
	for i, c := range cases {
		if got, want := Pow(c.x, c.y), c.want; got != want {
			t.Errorf("[%d]: Pow(%s,%s): got [%d](%s) want [%d](%s)",
				i,
				c.x,
				c.y,
				got,
				got,
				want,
				want,
			)
		}
	}
}

func TestPowSpecial(t *testing.T) {
	var test = func(t *testing.T, x, y, want FX) {
		if got, want := Pow(x, y), want; got != want {
			t.Errorf("Pow(%s,%s): got [%d](%s) want [%d](%s)",
				x,
				y,
				got,
				got,
				want,
				want,
			)
		}
	}
	test(t, Pi, NaN, NaN)
	test(t, NaN, Pi, NaN)
	test(t, One, Pi, One)
	test(t, -One, Inf, NaN)
	test(t, -One, -Inf, Zero)
	test(t, Zero, Zero, NaN)
	test(t, Zero, -Inf, NaN)
	test(t, Zero, Inf, Zero)
	test(t, Zero, Pi, Zero)
	test(t, Zero, Iota, Zero)
	test(t, Zero, -Iota, Zero)

	test(t, -Infs, Inf, Zero)
	test(t, -Infs, -Inf, NaN)
	test(t, -Infs, -Infs, NaN)
	test(t, -Infs, Infs, NaN)
	test(t, -Infs, -3*Pi, NaN)
	test(t, -Infs, 3*Pi, NaN)
	test(t, -Infs, Iota, NaN)
	test(t, -Infs, -Iota, NaN)
	test(t, -Infs, Zero, One)
	test(t, -Infs, I(18), Infs)
	test(t, -Infs, I(7531), -Infs)

	test(t, Infs, Inf, Zero)
	test(t, Infs, -Inf, Inf)
	test(t, Infs, -Infs, One)
	test(t, Infs, Infs, One)
	test(t, Infs, -3*Pi, Inf)
	test(t, Infs, 3*Pi, Infs)
	test(t, Infs, Iota, One)
	test(t, Infs, -Iota, One)

	test(t, -Iota, Inf, NaN)
	test(t, -Iota, -Inf, Zero)
	test(t, -Iota, -Infs, One)
	test(t, -Iota, Infs, One)
	test(t, -Iota, -3*Pi, NaN)
	test(t, -Iota, 3*Pi, NaN)
	test(t, Iota, Inf, Zero)
	test(t, Iota, -Inf, Zero)
	test(t, Iota, -Infs, One)
	test(t, Iota, Infs, One)
	test(t, Iota, -3*Pi, Inf)
	test(t, Iota, 3*Pi, Infs)

	test(t, -Inf, -3*Pi, 0)
	test(t, -Inf, 3*Pi, -Inf)
	test(t, -Inf, Zero, NaN)
	test(t, Inf, -3*Pi, 0)
	test(t, Inf, 3*Pi, Inf)
	test(t, Inf, Zero, NaN)

	test(t, Pi, Infs, One)
	test(t, Pi, Zero, One)
	test(t, Pi, -Infs, One)
	test(t, Pi, One, Pi)
	test(t, Pi, -Inf, Zero)
	test(t, Pi, Inf, Inf)

	test(t, -Pi, Infs, One)
	test(t, -Pi, Zero, One)
	test(t, -Pi, -Infs, One)
	test(t, -Pi, One, -Pi)
	test(t, -Pi, -Inf, Zero)
	test(t, -Pi, Inf, NaN)

	test(t, PiInv, Infs, One)
	test(t, PiInv, Zero, One)
	test(t, PiInv, -Infs, One)
	test(t, PiInv, One, PiInv)
	test(t, PiInv, -Inf, Zero)
	test(t, PiInv, Inf, Zero)

	test(t, -PiInv, Infs, One)
	test(t, -PiInv, Zero, One)
	test(t, -PiInv, -Infs, One)
	test(t, -PiInv, One, -PiInv)
	test(t, -PiInv, -Inf, Zero)
	test(t, -PiInv, Inf, NaN)
}
