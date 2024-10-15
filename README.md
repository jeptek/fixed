# Fixed point arithmetic in Go

[![codecov](https://codecov.io/gh/jeptek/fixed/branch/master/graph/badge.svg)](https://codecov.io/gh/jeptek/fixed)
[![Go Report Card](https://goreportcard.com/badge/github.com/jeptek/fixed)](https://goreportcard.com/report/github.com/jeptek/fixed)
[![Go Reference](https://pkg.go.dev/badge/github.com/jeptek/fixed?status.svg)](https://pkg.go.dev/github.com/jeptek/fixed?tab=doc)

`jeptek/fixed` is library for doing fixed-point arithmetic with Q31.32 numbers in Go.

**Supported operators are:**

- Add
- Sub
- Mul
- Div
- Pow
- Exp
- Ln
- Log2
- Sqrt


## Clone the project

```
$ git clone https://github.com/jeptek/fixed
$ cd fixed
```

## Using `jeptek/fixed`

```go
package main

import (
  "github.com/jeptek/fixed/q3132"
)

func main() {
  // (Just add negative sign to get negative numbers.)
  i17 := q3132.I(17)
  f3_235 := q3132.F(3.235)
  raw := q3132.FX(49263274885) // <== round_to_int(11.47*2^32) <= 11.47 

  fmt.Printf("%s = %s^%s",i17.Pow(f3_235), i17, f3_235) 
  fmt.Printf("%s = %s/%s",i17.Div(f3_235), i17, f3_235) 
  fmt.Printf("%s = %s*%s",i17.Mul(f3_235), i17, f3_235) 
  fmt.Printf("%s = %s+%s",i17.Add(f3_235), i17, f3_235) 
  fmt.Printf("%s = %s-%s",i17.Sub(f3_235), i17, f3_235) 
  fmt.Printf("%s = %s<%s",i17.Le(f3_235), i17, f3_235) 
  fmt.Printf("%s = %s<=%s",i17.Leq(f3_235), i17, f3_235) 
  fmt.Printf("%s = %s>%s",i17.Ge(f3_235), i17, f3_235) 
  fmt.Printf("%s = %s>=%s",i17.Geq(f3_235), i17, f3_235) 
  fmt.Printf("%s = sqrt(%s)",raw.Sqrt(), raw) 
  fmt.Printf("%s = exp(%s)",raw.Exp(), raw) 
  fmt.Printf("%s = ln(%s)",raw.Ln(), raw) 
  fmt.Printf("%s = log2(%s)",raw.Log2(), raw) 

  // you can also use functions directly for example q3132.Pow(x,y)

  var _floor int = raw.Floor()
  var _round int = raw.Round()
  var _ceil int = raw.Ceil()
  var _value float64 = raw.Value() // will return closest float64 number
}
```
