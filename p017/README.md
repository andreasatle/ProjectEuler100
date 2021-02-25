## Problem 17 - Number letter counts
If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.

If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?


NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23 letters and 115 (one hundred and fifteen) contains 20 letters. The use of "and" when writing out numbers is in compliance with British usage.

## Solution
Please see the git-repo for the code. There is nothing fancy here, but quite lengthy.
The LetterCountV1 creates the string for each number at the end, and LetterCountV2 counts
the letters as they come.

## Benchmark and test
Output:
```
✗ go test -bench=. -v
=== RUN   TestLetterCountV1
--- PASS: TestLetterCountV1 (0.00s)
=== RUN   TestLetterCountV2
--- PASS: TestLetterCountV2 (0.00s)
goos: darwin
goarch: amd64
pkg: github.com/andreasatle/ProjectEuler100/p017
BenchmarkLetterCountV1
BenchmarkLetterCountV1-12    	    9082	    131687 ns/op
BenchmarkLetterCountV2
BenchmarkLetterCountV2-12    	  110444	     12093 ns/op
PASS
ok  	github.com/andreasatle/ProjectEuler100/p017	3.618s
```

It is clear that LetterCountV2 is more than 10 times faster than LetterCountV1. Not very surprising.
One can clearly argue that writing out everything explicitly is clearer and less error prone, at the cost of the longer computational time.
