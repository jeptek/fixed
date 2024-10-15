## Benchmark setup

goos: linux
goarch: amd64
pkg: github.com/jeptek/fixed/q3132
cpu: Intel(R) Core(TM) i7-8086K CPU @ 4.00GHz


## Div
BenchmarkDiv_Pi_Pi                   1000000000               0.0000002 ns/op               0 B/op          0 allocs/op
BenchmarkDiv_Pi_PiInv                1000000000               0.0000002 ns/op               0 B/op          0 allocs/op
BenchmarkDiv_Pi_Two                  1000000000               0.0000003 ns/op               0 B/op          0 allocs/op
BenchmarkDiv_PiInv_Pi                1000000000               0.0000002 ns/op               0 B/op          0 allocs/op
BenchmarkDiv_PiInv_PiInv             1000000000               0.0000002 ns/op               0 B/op          0 allocs/op
BenchmarkDiv_PiInv_Two               1000000000               0.0000002 ns/op               0 B/op          0 allocs/op
BenchmarkDiv_Two_Two                 1000000000               0.0000002 ns/op               0 B/op          0 allocs/op
BenchmarkDiv_Two_Pi                  1000000000               0.0000003 ns/op               0 B/op          0 allocs/op
BenchmarkDiv_Two_PiInv               1000000000               0.0000002 ns/op               0 B/op          0 allocs/op
BenchmarkDiv_NeqTwo_Two              1000000000               0.0000002 ns/op               0 B/op          0 allocs/op
BenchmarkDiv_NeqTwo_Pi               1000000000               0.0000003 ns/op               0 B/op          0 allocs/op
BenchmarkDiv_NeqTwo_PiInv            1000000000               0.0000003 ns/op               0 B/op          0 allocs/op
BenchmarkDiv_NeqTwo_NeqTwo           1000000000               0.0000002 ns/op               0 B/op          0 allocs/op
BenchmarkDiv_NeqTwo_NeqPi            1000000000               0.0000003 ns/op               0 B/op          0 allocs/op
BenchmarkDiv_NeqTwo_NeqPiInv         1000000000               0.0000003 ns/op               0 B/op          0 allocs/op

## Exp
BenchmarkExp_Pi                      1000000000               0.0000002 ns/op               0 B/op          0 allocs/op
BenchmarkExp_PiInv                   1000000000               0.0000003 ns/op               0 B/op          0 allocs/op
BenchmarkExp_Two                     1000000000               0.0000003 ns/op               0 B/op          0 allocs/op
BenchmarkExp_Four                    1000000000               0.0000003 ns/op               0 B/op          0 allocs/op

## Ln
BenchmarkLn_Pi                       1000000000               0.0000007 ns/op               0 B/op          0 allocs/op
BenchmarkLn_PiInv                    1000000000               0.0000010 ns/op               0 B/op          0 allocs/op
BenchmarkLn_Two                      1000000000               0.0000007 ns/op               0 B/op          0 allocs/op
BenchmarkLn_Four                     1000000000               0.0000007 ns/op               0 B/op          0 allocs/op

## Log2
BenchmarkLog2_Pi                     1000000000               0.0000004 ns/op               0 B/op          0 allocs/op
BenchmarkLog2_PiInv                  1000000000               0.0000005 ns/op               0 B/op          0 allocs/op
BenchmarkLog2_Two                    1000000000               0.0000002 ns/op               0 B/op          0 allocs/op
BenchmarkLog2_Four                   1000000000               0.0000002 ns/op               0 B/op          0 allocs/op

## Mul
BenchmarkMul_Pi_Pi                   1000000000               0.0000002 ns/op               0 B/op          0 allocs/op
BenchmarkMul_Pi_PiInv                1000000000               0.0000002 ns/op               0 B/op          0 allocs/op
BenchmarkMul_Pi_Two                  1000000000               0.0000002 ns/op               0 B/op          0 allocs/op
BenchmarkMul_PiInv_Pi                1000000000               0.0000002 ns/op               0 B/op          0 allocs/op
BenchmarkMul_PiInv_PiInv             1000000000               0.0000002 ns/op               0 B/op          0 allocs/op
BenchmarkMul_PiInv_Two               1000000000               0.0000002 ns/op               0 B/op          0 allocs/op
BenchmarkMul_Two_Two                 1000000000               0.0000002 ns/op               0 B/op          0 allocs/op
BenchmarkMul_Two_Pi                  1000000000               0.0000002 ns/op               0 B/op          0 allocs/op
BenchmarkMul_Two_PiInv               1000000000               0.0000002 ns/op               0 B/op          0 allocs/op
BenchmarkMul_NeqTwo_Two              1000000000               0.0000002 ns/op               0 B/op          0 allocs/op
BenchmarkMul_NeqTwo_Pi               1000000000               0.0000003 ns/op               0 B/op          0 allocs/op
BenchmarkMul_NeqTwo_PiInv            1000000000               0.0000002 ns/op               0 B/op          0 allocs/op
BenchmarkMul_NeqTwo_NeqTwo           1000000000               0.0000002 ns/op               0 B/op          0 allocs/op
BenchmarkMul_NeqTwo_NeqPi            1000000000               0.0000002 ns/op               0 B/op          0 allocs/op
BenchmarkMul_NeqTwo_NeqPiInv         1000000000               0.0000002 ns/op               0 B/op          0 allocs/op

## Pow2
BenchmarkPow2_Pi                     1000000000               0.0000006 ns/op               0 B/op          0 allocs/op
BenchmarkPow2_NeqPi                  1000000000               0.0000005 ns/op               0 B/op          0 allocs/op
BenchmarkPow2_PiInv                  1000000000               0.0000004 ns/op               0 B/op          0 allocs/op
BenchmarkPow2_NeqPiInv               1000000000               0.0000005 ns/op               0 B/op          0 allocs/op
BenchmarkPow2_Two                    1000000000               0.0000003 ns/op               0 B/op          0 allocs/op
BenchmarkPow2_NeqTwo                 1000000000               0.0000003 ns/op               0 B/op          0 allocs/op
BenchmarkPow2_Four                   1000000000               0.0000003 ns/op               0 B/op          0 allocs/op
BenchmarkPow2_NeqFour                1000000000               0.0000004 ns/op               0 B/op          0 allocs/op

## Pow
BenchmarkPow_Pi_Pi                   1000000000               0.0000010 ns/op               0 B/op          0 allocs/op
BenchmarkPow_Pi_PiInv                1000000000               0.0000015 ns/op               0 B/op          0 allocs/op
BenchmarkPow_Pi_Two                  1000000000               0.0000008 ns/op               0 B/op          0 allocs/op
BenchmarkPow_PiInv_Pi                1000000000               0.0000008 ns/op               0 B/op          0 allocs/op
BenchmarkPow_PiInv_PiInv             1000000000               0.0000012 ns/op               0 B/op          0 allocs/op
BenchmarkPow_PiInv_Two               1000000000               0.0000008 ns/op               0 B/op          0 allocs/op
BenchmarkPow_Two_Two                 1000000000               0.0000004 ns/op               0 B/op          0 allocs/op
BenchmarkPow_Two_Pi                  1000000000               0.0000005 ns/op               0 B/op          0 allocs/op
BenchmarkPow_Two_PiInv               1000000000               0.0000006 ns/op               0 B/op          0 allocs/op
BenchmarkPow_NeqTwo_Two              1000000000               0.0000003 ns/op               0 B/op          0 allocs/op
BenchmarkPow_NeqTwo_Pi               1000000000               0.0000002 ns/op               0 B/op          0 allocs/op
BenchmarkPow_NeqTwo_NeqTwo           1000000000               0.0000004 ns/op               0 B/op          0 allocs/op
BenchmarkPow_NeqTwo_100              1000000000               0.0000002 ns/op               0 B/op          0 allocs/op
BenchmarkPow_NeqTwo_101              1000000000               0.0000002 ns/op               0 B/op          0 allocs/op
BenchmarkPow_NeqTwo_Neq100           1000000000               0.0000002 ns/op               0 B/op          0 allocs/op
BenchmarkPow_NeqTwo_Neq101           1000000000               0.0000002 ns/op               0 B/op          0 allocs/op

## Sqrt
BenchmarkSqrt_Pi                     1000000000               0.0000003 ns/op               0 B/op          0 allocs/op
BenchmarkSqrt_PiInv                  1000000000               0.0000002 ns/op               0 B/op          0 allocs/op
BenchmarkSqrt_Two                    1000000000               0.0000003 ns/op               0 B/op          0 allocs/op
BenchmarkSqrt_Four                   1000000000               0.0000003 ns/op               0 B/op          0 allocs/op