package q3132

import "testing"

func TestAdd(t *testing.T) {
	type op struct {
		x    FX
		y    FX
		want FX
	}

	cases := []op{
		{x: -PiInv, y: PiInv, want: Zero},
		{x: PiInv, y: -PiInv, want: Zero},
		{x: -PiInv, y: -PiInv, want: -2 * PiInv},
		{x: MaxValue - One, y: MaxValue - I(423543), want: Inf},        // overflow
		{x: -MaxValue + One, y: -MaxValue + I(423543), want: -Inf},     // underflow
		{x: F(1234.5), y: F(-8888.875), want: -32875290296320},         //-7654.375000000000000
		{x: F(1234.5), y: F(8888.875), want: 43479564550144},           //10123.375000000000000
		{x: F(-14423.35375565), y: I(-17), want: -62020847123187},      //-14440.353755649877712
		{x: F(-14423.35375565), y: I(-14), want: -62007962221299},      //-14437.353755649877712
		{x: 4423816314, y: I(727), want: 3126865040506},                //728.029999999795109
		{x: F(-1.35375565), y: I(-3), want: -18699238131},              //-4.353755649877712
		{x: F(-1.35375565), y: I(5), want: 15660500237},                //3.646244350122288
		{x: F(-1.35375565), y: I(12), want: 45725271309},               //10.646244350122288
		{x: F(11.17823115), y: F(-5.8754541), want: 22775254008},       //5.302777050063014
		{x: F(13.33761137), y: F(5.87523451), want: 82518544716},       //19.212845879606903
		{x: F(0.015625), y: F(0.000244140625), want: 68157440},         //0.015869140625000
		{x: F(0.500244140625), y: F(0.500732421875), want: 4299161600}, //1.000976562500000
		{x: F(1.25), y: F(1.5), want: 11811160064},                     //2.750000000000000
		{x: F(1.25), y: I(-4), want: -11811160064},                     //-2.750000000000000
		{x: F(1.25), y: I(4), want: 22548578304},                       //5.250000000000000
		{x: F(1.44140625), y: F(1.44140625), want: 12381585408},        //2.882812500000000
		{x: F(1.44140625), y: F(1.441650390625), want: 12382633984},    //2.883056640625000
		{x: F(1.515625), y: F(1.531250), want: 13086228480},            //3.046875000000000
		{x: Pi, y: E, want: 25167969258},                               //5.859874481800944
		{x: Pi, y: -E, want: 1818106150},                               //-0.423310825135559
		{x: E, y: Pi, want: 25167969258},                               //5.859874481800944
		{x: E, y: -Pi, want: -1818106150},                              //-0.423310825135559
		{x: E, y: PiInv, want: 13042062105},                            //3.036591714480892
		{x: E, y: -PiInv, want: 10307801003},                           //2.399971942184493
		{x: PiInv, y: PiInv, want: 2734261102},                         //0.636619772296399

	}
	for i, c := range cases {
		if got, want := Add(c.x, c.y), c.want; got != want {
			t.Errorf("[%d]: Add(%s,%s): got [%d](%s) want [%d](%s)",
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

func TestAddSpecial(t *testing.T) {
	var test = func(t *testing.T, x, y, want FX) {
		if got, want := Add(x, y), want; got != want {
			t.Errorf("Add(%s,%s): got [%d](%s) want [%d](%s)",
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

	test(t, One, Pi, 17788005000) // 4.141592653468251
	test(t, -One, Inf, Inf)
	test(t, -One, -Inf, -Inf)

	test(t, Zero, Zero, Zero)
	test(t, Zero, -Inf, -Inf)
	test(t, Zero, Inf, Inf)
	test(t, Zero, -Infs, -Infs)
	test(t, Zero, Infs, Infs)
	test(t, Zero, Pi, Pi)
	test(t, Zero, Iota, Iota)
	test(t, Zero, -Iota, -Iota)

	test(t, -Infs, Inf, Inf)
	test(t, -Infs, -Inf, -Inf)
	test(t, -Infs, -Infs, -Infs)
	test(t, -Infs, Infs, Zero)
	test(t, -Infs, -3*Pi, -3*Pi)
	test(t, -Infs, 3*Pi, 3*Pi)
	test(t, -Infs, Iota, Iota)
	test(t, -Infs, -Iota, -Iota)
	test(t, -Infs, Zero, -Infs)
	test(t, -Infs, I(18), I(18))
	test(t, -Infs, I(7531), I(7531))

	test(t, Infs, Zero, Infs)
	test(t, Infs, Inf, Inf)
	test(t, Infs, -Inf, -Inf)
	test(t, Infs, -Infs, Zero)
	test(t, Infs, Infs, Infs)
	test(t, Infs, -3*Pi, -3*Pi)
	test(t, Infs, 3*Pi, 3*Pi)
	test(t, Infs, Iota, Iota)
	test(t, Infs, -Iota, -Iota)

	test(t, -Iota, Inf, Inf)
	test(t, -Iota, -Inf, -Inf)
	test(t, -Iota, -Infs, -Iota)
	test(t, -Iota, Infs, -Iota)
	test(t, -Iota, -3*Pi, -40479113113) //-9.424777960637584
	test(t, -Iota, 3*Pi, 40479113111)   //9.424777960171923

	test(t, Iota, Inf, Inf)
	test(t, Iota, -Inf, -Inf)
	test(t, Iota, -Infs, Iota)
	test(t, Iota, Infs, Iota)
	test(t, Iota, -3*Pi, -40479113111) //-9.424777960171923
	test(t, Iota, 3*Pi, 40479113113)   //9.424777960637584

	test(t, -Inf, -Inf, -Inf)
	test(t, -Inf, +Inf, NaN)
	test(t, -Inf, Zero, -Inf)
	test(t, -Inf, -Infs, -Inf)
	test(t, -Inf, +Infs, -Inf)
	test(t, -Inf, -3*Pi, -Inf)
	test(t, -Inf, 3*Pi, -Inf)

	test(t, Inf, -Inf, NaN)
	test(t, Inf, Inf, Inf)
	test(t, Inf, Zero, Inf)
	test(t, Inf, -Infs, Inf)
	test(t, Inf, +Infs, Inf)
	test(t, Inf, -3*Pi, Inf)
	test(t, Inf, 3*Pi, Inf)

	test(t, Pi, Infs, Pi)
	test(t, Pi, Zero, Pi)
	test(t, Pi, -Infs, Pi)
	test(t, Pi, One, 17788005000) //4.141592653468251
	test(t, Pi, -Inf, -Inf)
	test(t, Pi, Inf, Inf)

	test(t, -Pi, Infs, -Pi)
	test(t, -Pi, Zero, -Pi)
	test(t, -Pi, -Infs, -Pi)
	test(t, -Pi, One, -9198070408) // -2.141592653468251
	test(t, -Pi, -Inf, -Inf)
	test(t, -Pi, Inf, Inf)

	test(t, PiInv, Infs, PiInv)
	test(t, PiInv, Zero, PiInv)
	test(t, PiInv, -Infs, PiInv)
	test(t, PiInv, One, 5662097847) //1.318309886148199
	test(t, PiInv, -Inf, -Inf)
	test(t, PiInv, Inf, Inf)

	test(t, -PiInv, Infs, -PiInv)
	test(t, -PiInv, Zero, -PiInv)
	test(t, -PiInv, -Infs, -PiInv)
	test(t, -PiInv, One, 2927836745) // 0.681690113851801
	test(t, -PiInv, -Inf, -Inf)
	test(t, -PiInv, Inf, Inf)
}
