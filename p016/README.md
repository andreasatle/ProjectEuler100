## Problem 16 - Power digit sum
2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2^1000?

## Solution
I utilize the math/big package from golang, to handle big numbers.
The code computes the (very large) number base^pow, and we extract 
the 18 last numbers and convert them to int64 before computing the
digit sum. I assume that it is faster to use integer arithmetics than
the big-numbers in golang, so it should be faster to convert as large chunk
of the number as possible every time. The big-package is a bit awkward to
fully comprehend. There are a lot of work arrays, that need to be allocated
ahead of time. It works though...
Code:
```go
// PowerDigitSum returns the digit sum of base^pow.
func PowerDigitSum(base, pow int) int {
	zeroBig := &big.Int{}
	baseBig := big.NewInt(int64(base))
	powBig := big.NewInt(int64(pow))
	splitBig := big.NewInt(int64(1000000000000000000))
	zBig := &big.Int{}
	mBig := &big.Int{}
    zBig = zBig.Exp(baseBig, powBig, mBig)

	digitSum := 0
	for zBig.Cmp(zeroBig) == 1 {
		zBig.DivMod(zBig, splitBig, mBig)
		digitSum += computeDigitSum(mBig.Int64())
	}

	return digitSum
}

func computeDigitSum(n int64) int {
	nn := int(n)
	sum := 0
	for ; nn > 0; nn /= 10 {
		sum += nn % 10
	}
	return sum
}
```

## Benchmark and test
Code:
```go
func TestPowerDigitSum(t *testing.T) {
	assert.Equal(t, 26, p016.PowerDigitSum(2, 15))
	assert.Equal(t, 1366, p016.PowerDigitSum(2, 1000))
}

func BenchmarkPowerDigitSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p016.PowerDigitSum(2, 1000)
	}
}
```
with output:
```
✗ go test -bench=. -v
=== RUN   TestPowerDigitSum
--- PASS: TestPowerDigitSum (0.00s)
goos: darwin
goarch: amd64
pkg: github.com/andreasatle/ProjectEuler100/p016
BenchmarkPowerDigitSum
BenchmarkPowerDigitSum-12    	  225562	      4562 ns/op
PASS
ok  	github.com/andreasatle/ProjectEuler100/p016	1.099s
```

The code takes less than 5 micro-seconds to complete.