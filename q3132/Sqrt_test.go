package q3132

import (
	"testing"
)

func BenchmarkSqrt_Pi(b *testing.B) {
	Sqrt(Pi)
}

func BenchmarkSqrt_PiInv(b *testing.B) {
	Sqrt(PiInv)
}

func BenchmarkSqrt_Two(b *testing.B) {
	Sqrt(Two)
}

func BenchmarkSqrt_Four(b *testing.B) {
	const four = 4 * One
	Sqrt(four)
}

func TestSqrt(t *testing.T) {
	type op struct {
		in   FX
		want FX
	}

	cases := []op{
		{in: 9223372036854775805, want: 199032864766655},
		{in: 922337203685477580, want: 62939718189019},
		{in: 92233720368547758, want: 19903286525570},
		{in: 9223372036854776, want: 6293971818906},
		{in: 922337203685478, want: 1990328647664},
		{in: 92233720368548, want: 629397183014},
		{in: 9223372036855, want: 199032864766},
		{in: 922337203685, want: 62939718189},
		{in: 92233720369, want: 19903286502},
		{in: 9223372037, want: 6293971819},
		{in: 922337204, want: 1990328648},
		{in: 92233720, want: 629397180},
		{in: 9223372, want: 199032864},
		{in: 922337, want: 62939711},
		{in: 92234, want: 19903316},
		{in: 9223, want: 6293845},
		{in: 922, want: 1989965},
		{in: 92, want: 628599},
		{in: 9, want: 196608},
		{in: F(1234.5), want: 150905618855},     //35.135452368995175
		{in: F(1.35375565), want: 4997237644},   //1.163510057143867
		{in: F(11.17823115), want: 14359734302}, //3.343386180233210
		{in: F(13.33761137), want: 15685518910}, //3.652069463860244
		{in: F(0.015625), want: F(0.125)},
		{in: F(0.500244140625), want: 3037741864}, //0.706920669
		{in: F(1.25), want: 4801919433},           //1.118033992359415
		{in: F(1.44140625), want: 5156476726},     //1.200585795100778
		{in: F(1.515625), want: 5287565271},       //1.231107225408778
		{in: Pi, want: 7612631323},                //1.772453850833699
		{in: E, want: 7081203937},                 //1.648721270496026
		{in: PiInv, want: 2423175810},             //0.564189583528787

	}
	for i, c := range cases {
		if got, want := Sqrt(c.in), c.want; got != want {
			t.Errorf("[%d]: Sqrt(%s): got [%d](%s) want [%d](%s) >>> diff: %.15f%%",
				i,
				c.in,
				got,
				got,
				want,
				want,
				100*float64(got-want)/float64(want),
			)
		}
	}
}

func TestSqrtSpecial(t *testing.T) {
	var test = func(t *testing.T, x, want FX) {
		if got, want := Sqrt(x), want; got != want {
			t.Errorf("Sqrt([%d]%s): got [%d](%s) want [%d](%s)",
				x,
				x,
				got,
				got,
				want,
				want,
			)
		}
	}

	test(t, NaN, NaN)
	test(t, -Inf, NaN)
	test(t, -Infs, NaN)
	test(t, -One, NaN)
	test(t, Inf, Inf)
	test(t, Infs, Infs)
	test(t, One, One)
	test(t, Zero, Zero)
	test(t, Iota, 65536) //0.000015258789062
}
