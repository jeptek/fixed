package q3132

import (
	"math"
	"testing"
)

func BenchmarkMul_Pi_Pi(b *testing.B) {
	Mul(Pi, Pi)
}
func BenchmarkMul_Pi_PiInv(b *testing.B) {
	Mul(Pi, PiInv)
}

func BenchmarkMul_Pi_Two(b *testing.B) {
	Mul(Pi, Two)
}

func BenchmarkMul_PiInv_Pi(b *testing.B) {
	Mul(PiInv, Pi)
}
func BenchmarkMul_PiInv_PiInv(b *testing.B) {
	Mul(PiInv, PiInv)
}

func BenchmarkMul_PiInv_Two(b *testing.B) {
	Mul(PiInv, Two)
}

func BenchmarkMul_Two_Two(b *testing.B) {
	Mul(Two, Two)
}

func BenchmarkMul_Two_Pi(b *testing.B) {
	Mul(Two, Pi)
}

func BenchmarkMul_Two_PiInv(b *testing.B) {
	Mul(Two, PiInv)
}

func BenchmarkMul_NeqTwo_Two(b *testing.B) {
	Mul(-Two, Two)
}

func BenchmarkMul_NeqTwo_Pi(b *testing.B) {
	Mul(-Two, Pi)
}

func BenchmarkMul_NeqTwo_PiInv(b *testing.B) {
	Mul(-Two, PiInv)
}

func BenchmarkMul_NeqTwo_NeqTwo(b *testing.B) {
	Mul(-Two, -Two)
}

func BenchmarkMul_NeqTwo_NeqPi(b *testing.B) {
	Mul(-Two, -Pi)
}

func BenchmarkMul_NeqTwo_NeqPiInv(b *testing.B) {
	Mul(-Two, -PiInv)
}

func TestMul(t *testing.T) {

	type op struct {
		x    FX
		y    FX
		want FX
	}

	cases := []op{
		{x: 5115377064673848749, y: 5115377064673848750, want: Inf},
		{x: One, y: 6203371737317376787, want: 6203371737317376787},
		{x: One, y: 6203371737317376, want: 6203371737317376},
		{x: F(1.25), y: I(-4), want: I(-5)},
		{x: F(-1.25), y: I(4), want: I(-5)},
		{x: F(-1.25), y: I(-4), want: I(5)},
		{x: F(1.25), y: I(4), want: I(5)},
		{x: E, y: -Pi, want: -36677879200},                             //-8.539734221994877
		{x: E, y: -PiInv, want: -3716246134},                           //-0.865255979355425
		{x: E, y: Pi, want: 36677879200},                               //8.539734221994877
		{x: E, y: PiInv, want: 3716246134},                             //0.865255979355425
		{x: F(0.015625), y: F(0.000244140625), want: 16384},            //0.000003814697266
		{x: F(0.500244140625), y: F(0.500732421875), want: 1075839744}, //0.250488460063934
		{x: F(1.25), y: F(1.5), want: 8053063680},                      //1.875000000000000
		{x: F(1.44140625), y: F(1.44140625), want: 8923447296},         //2.077651977539062
		{x: F(1.44140625), y: F(1.441650390625), want: 8924958720},     //2.078003883361816
		{x: F(1.515625), y: F(1.531250), want: 9967763456},             //2.320800781250000
		{x: F(11.17823115), y: F(-5.8754541), want: -282081357540},     //-65.677184038795531
		{x: F(1234.5), y: F(-8888.875), want: -47130034153979904},      //-10973316.187500000000000
		{x: F(1234.5), y: F(8888.875), want: 47130034153979904},        // 10973316.187500000000000
		{x: F(13.33761137), y: F(5.87523451), want: 336560486062},      // 78.361594598274678
		{x: Pi, y: -E, want: -36677879200},                             //-8.539734221994877
		{x: Pi, y: E, want: 36677879200},                               //8.539734221994877
		{x: Pi, y: PiInv, want: 4294967295},                            //0.999999999767169
		{x: PiInv, y: -PiInv, want: -435171170},                        //-0.101321183610708
		{x: PiInv, y: PiInv, want: 435171170},                          // 0.101321183610708
	}
	for i, c := range cases {
		if got, want := Mul(c.x, c.y), c.want; got != want {
			t.Errorf("\n\t\t[%d]: Mul(%d,%d): got %d want %d  >>> diff: %.15f%%\n\t\t\t%s*%s=%s : got: %s",
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

func TestMulSpecial(t *testing.T) {
	var test = func(t *testing.T, x, y, want FX) {
		if got, want := Mul(x, y), want; got != want {
			t.Errorf("\n\t\tMul(%d,%d): got %d want %d  >>> diff: %.15f%%\n\t\t\t%s*%s=%s : got: %s",
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
	test(t, One, Pi, Pi)

	test(t, Zero, Zero, Zero)
	test(t, Zero, -Inf, Zero)
	test(t, Zero, Inf, Zero)
	test(t, Zero, Pi, Zero)

	test(t, -Infs, Inf, -Inf)
	test(t, -Infs, -Inf, Inf)
	test(t, -Infs, -Infs, Infs)
	test(t, -Infs, Infs, -Infs)
	test(t, -Infs, -3*Pi, Infs)
	test(t, -Infs, 3*Pi, -Infs)

	test(t, Infs, Inf, Inf)
	test(t, Infs, -Inf, -Inf)
	test(t, Infs, -Infs, -Infs)
	test(t, Infs, Infs, Infs)
	test(t, Infs, -3*Pi, -Infs)
	test(t, Infs, 3*Pi, Infs)

	test(t, -Inf, -3*Pi, Inf)
	test(t, -Inf, 3*Pi, -Inf)
	test(t, -Inf, Zero, Zero)

	test(t, Inf, -3*Pi, -Inf)
	test(t, Inf, 3*Pi, Inf)
	test(t, Inf, Zero, Zero)

	test(t, Iota, Iota, Infs)
	test(t, Iota, -Iota, -Infs)
	test(t, Iota, Inf, Inf)
	test(t, Iota, -Inf, -Inf)
	test(t, Iota, MaxValue, Half)
	test(t, Iota, -MaxValue, -Half)

	test(t, -Iota, Iota, -Infs)
	test(t, -Iota, -Iota, Infs)
	test(t, -Iota, Inf, -Inf)
	test(t, -Iota, -Inf, Inf)
	test(t, -Iota, MaxValue, -Half)
	test(t, -Iota, -MaxValue, Half)

	test(t, MaxValue, Iota, Half)
	test(t, MaxValue, -Iota, -Half)
	test(t, MaxValue, Inf, Inf)
	test(t, MaxValue, -Inf, -Inf)
	test(t, MaxValue, MaxValue, Inf)
	test(t, MaxValue, -MaxValue, -Inf)

	test(t, -MaxValue, Iota, -Half)
	test(t, -MaxValue, -Iota, Half)
	test(t, -MaxValue, Inf, -Inf)
	test(t, -MaxValue, -Inf, Inf)
	test(t, -MaxValue, MaxValue, -Inf)
	test(t, -MaxValue, -MaxValue, Inf)

	test(t, -Iota, Iota, -Infs)
	test(t, -Iota, -Iota, Infs)
	test(t, -Iota, Inf, -Inf)
	test(t, -Iota, -Inf, Inf)
	test(t, -Iota, MaxValue, -Half)
	test(t, -Iota, -MaxValue, Half)

	test(t, Pi, Infs, Infs)
	test(t, Pi, Zero, Zero)
	test(t, Pi, -Infs, -Infs)
	test(t, Pi, One, Pi)
	test(t, Pi, -Inf, -Inf)
	test(t, Pi, Inf, Inf)

	test(t, -Pi, Infs, -Infs)
	test(t, -Pi, Zero, Zero)
	test(t, -Pi, -Infs, Infs)
	test(t, -Pi, One, -Pi)
	test(t, -Pi, -Inf, Inf)
	test(t, -Pi, Inf, -Inf)
}

func TestMulByOneMinusIota(t *testing.T) {
	const (
		totalBits = 63
		fracBits  = 32

		oneMinusIota  = FX(1<<fracBits) - 1
		oneMinusIotaF = float64(oneMinusIota) / (1 << fracBits)
	)

	for _, neg := range []bool{false, true} {
		for i := uint(0); i < totalBits; i++ {
			x := FX(1 << i)
			if neg {
				x = -x
			} else if i == totalBits-1 {
				// A signed int64 can't represent 1<<63.
				continue
			}

			// want equals x * oneMinusIota, rounded to nearest.
			want := FX(0)
			if -1<<fracBits < x && x < 1<<fracBits {
				// (x * oneMinusIota) isn't exactly representable as an
				// FX. Calculate the rounded value using float64 math.
				xF := float64(x) / (1 << fracBits)
				wantF := xF * oneMinusIotaF * (1 << fracBits)
				want = FX(math.Floor(wantF + 0.5))
			} else {
				// (x * oneMinusIota) is exactly representable.
				want = oneMinusIota << (i - fracBits)
				if neg {
					want = -want
				}
			}

			if got := Mul(x, oneMinusIota); got != want {
				t.Errorf("neg=%t, i=%d, x=%v, Mul: got %v, want %v", neg, i, x, got, want)
			}
		}
	}
}
