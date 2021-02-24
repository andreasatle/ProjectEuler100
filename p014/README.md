## Problem 14 - Longest Collatz sequence
The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:

13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.

## Solution
I programmed two versions. The first version loops over all numbers and compute the full sequence for each number. The second version aborts the sequence when an already computed value occurs. The second version start with the largest value, and ends with the smallest.
```go
// MaxCollatzSequenceV1 returns the number (not exceeding n) with the longest Collatz-sequence.
func MaxCollatzSequenceV1(n int) int {
	maxNum := 0
	maxCount := 0
	for num := 2; num <= n; num++ {
		cnt := 0
		for p := num; p > 1; p = nextCollatz(p) {
			cnt++
		}
		if cnt > maxCount {
			maxNum = num
			maxCount = cnt
		}
	}
	return maxNum
}

// MaxCollatzSequenceV2 returns the number (not exceeding n) with the longest Collatz-sequence.
func MaxCollatzSequenceV2(n int) int {
	seqLength := make([]int, n+1)

// Set the length starting at 2 to 1, so that algorithm terminates.
	seqLength[2] = 1

    // Assume that no sequence is larger than 1000 steps (not really necessary...)
	seq := make([]int, 0, 1000)
	for num := n; num > 1; num-- {
		// skip if num already computed
		if seqLength[num] > 0 {
			continue
		}
		// Compute next sequence until computed value found
		seq = seq[:0]
		for p := num; ; p = nextCollatz(p) {
			seq = append(seq, p)
			if p <= n && seqLength[p] > 0 {
				break
			}
		}
		lastSeq := len(seq) - 1
		for i, cnt := lastSeq, seqLength[seq[lastSeq]]; i >= 0; i, cnt = i-1, cnt+1 {
			if seq[i] <= n {
				seqLength[seq[i]] = cnt
			}
		}
	}
	return findIndexLargestValue(seqLength)
}

func nextCollatz(n int) int {
	if n%2 == 0 {
		return n / 2
	}
	return 3*n + 1
}

func findIndexLargestValue(slice []int) int {
	maxIndex := 0
	maxValue := slice[0]
	for i, v := range slice {
		if v > maxValue {
			maxValue = v
			maxIndex = i
		}
	}
	return maxIndex
}
```

## Test and Benchmark
Code:
```go
func TestMaxCollatzSequenceV1(t *testing.T) {
	assert.Equal(t, 837799, p014.MaxCollatzSequenceV1(1000000))
}

func TestMaxCollatzSequenceV2(t *testing.T) {
	assert.Equal(t, 837799, p014.MaxCollatzSequenceV2(1000000))
}

func BenchmarkMaxCollatzSequenceV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p014.MaxCollatzSequenceV1(1000000)
	}
}

func BenchmarkMaxCollatzSequenceV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p014.MaxCollatzSequenceV2(1000000)
	}
}
```
with output:
```
✗ go test -bench=. -v
=== RUN   TestMaxCollatzSequenceV1
--- PASS: TestMaxCollatzSequenceV1 (0.23s)
=== RUN   TestMaxCollatzSequenceV2
--- PASS: TestMaxCollatzSequenceV2 (0.02s)
goos: darwin
goarch: amd64
pkg: github.com/andreasatle/ProjectEuler100/p014
BenchmarkMaxCollatzSequenceV1
BenchmarkMaxCollatzSequenceV1-12    	       5	 222635921 ns/op
BenchmarkMaxCollatzSequenceV2
BenchmarkMaxCollatzSequenceV2-12    	      81	  16338822 ns/op
PASS
ok  	github.com/andreasatle/ProjectEuler100/p014	3.833s
```
The optimized version (V2) is about 13.6 times faster than the naive algorithm (V1) that computes the sequence for each number separately.

I haven't made any effort to make a concurrent algorithm. The first algorithm should be very easy to use go-routines. Then second use slices to store intermediate values, and are quite sequential in nature. It is not clear that a concurrent version would improve much.

## Attempt to use go-routines
V3 is a go-routine version of V1 and V4 is a go-routine version of V2.

Code:
```go
func MaxCollatzSequenceV3(n int) int {
	var wg sync.WaitGroup
	var mu sync.Mutex

	maxNum := 0
	maxCount := 0

	numProc := runtime.NumCPU()
	wg.Add(numProc)
	for iProc := 0; iProc < numProc; iProc++ {
		go func(iProc int, maxNum, maxCount *int, wg *sync.WaitGroup, mu *sync.Mutex) {
			for num := 2 + iProc; num <= n; num += numProc {
				cnt := 0
				for p := num; p > 1; p = nextCollatz(p) {
					cnt++
				}
				if cnt > *maxCount {
					mu.Lock()
					*maxNum = num
					*maxCount = cnt
					mu.Unlock()
				}
			}
			wg.Done()
		}(iProc, &maxNum, &maxCount, &wg, &mu)
	}
	wg.Wait()
	return maxNum
}

// MaxCollatzSequenceV4 returns the number (not exceeding n) with the longest Collatz-sequence.
func MaxCollatzSequenceV4(n int) int {
	var wg sync.WaitGroup
	//var mus []sync.Mutex
	//var mu sync.Mutex

	seqLength := make([]int, n+1)
	mus := make([]sync.Mutex, n+1)
	// Set the length starting at 2 to 1, so that algorithm terminates.
	seqLength[2] = 1

	numProc := runtime.NumCPU()
	wg.Add(numProc)

	for iProc := 0; iProc < numProc; iProc++ {

		go func(iProc int, wg *sync.WaitGroup, mus []sync.Mutex) {
			// Assume that no sequence is larger than 1000 steps (not really necessary...)
			seq := make([]int, 0, 1000)
			for num := n - iProc; num > 1; num -= numProc {
				// skip if num already computed
				if seqLength[num] > 0 {
					continue
				}
				// Compute next sequence until computed value found
				seq = seq[:0]
				for p := num; ; p = nextCollatz(p) {
					seq = append(seq, p)
					if p <= n && seqLength[p] > 0 {
						break
					}
				}
				lastSeq := len(seq) - 1
				for i, cnt := lastSeq, seqLength[seq[lastSeq]]; i >= 0; i, cnt = i-1, cnt+1 {
					if seq[i] <= n {
						mus[seq[i]].Lock()
						seqLength[seq[i]] = cnt
						mus[seq[i]].Unlock()
					}
				}
			}
			wg.Done()
		}(iProc, &wg, mus)
	}
	wg.Wait()
	return findIndexLargestValueV2(seqLength)
}

func nextCollatz(n int) int {
	if n%2 == 0 {
		return n / 2
	}
	return 3*n + 1
}

func findIndexLargestValue(slice []int) int {
	maxIndex := 0
	maxValue := slice[0]
	for i, v := range slice {
		if v > maxValue {
			maxValue = v
			maxIndex = i
		}
	}
	return maxIndex
}

func findIndexLargestValueV2(slice []int) int {
	var wg sync.WaitGroup
	var mu sync.Mutex

	maxIndex := 0
	maxValue := slice[0]

	numProc := runtime.NumCPU()
	wg.Add(numProc)

	for iProc := 0; iProc < numProc; iProc++ {
		go func(iProc int, maxValue, maxIndex *int, wg *sync.WaitGroup, mu *sync.Mutex) {
			for i, v := range slice {
				if v > *maxValue {
					mu.Lock()
					*maxValue = v
					*maxIndex = i
					mu.Unlock()
				}
			}
			wg.Done()
		}(iProc, &maxValue, &maxIndex, &wg, &mu)
	}

	wg.Wait()
	return maxIndex
}
```

Extended test and benchmark:
```go
func TestMaxCollatzSequenceV1(t *testing.T) {
	assert.Equal(t, 837799, p014.MaxCollatzSequenceV1(1000000))
}

func TestMaxCollatzSequenceV2(t *testing.T) {
	assert.Equal(t, 837799, p014.MaxCollatzSequenceV2(1000000))
}

func TestMaxCollatzSequenceV3(t *testing.T) {
	assert.Equal(t, 837799, p014.MaxCollatzSequenceV3(1000000))
}

func TestMaxCollatzSequenceV4(t *testing.T) {
	assert.Equal(t, 837799, p014.MaxCollatzSequenceV4(1000000))
}

func BenchmarkMaxCollatzSequenceV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p014.MaxCollatzSequenceV1(1000000)
	}
}

func BenchmarkMaxCollatzSequenceV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p014.MaxCollatzSequenceV2(1000000)
	}
}

func BenchmarkMaxCollatzSequenceV3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p014.MaxCollatzSequenceV3(1000000)
	}
}

func BenchmarkMaxCollatzSequenceV4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p014.MaxCollatzSequenceV4(1000000)
	}
}
```
with output:
```
✗ go test -bench=. -v
=== RUN   TestMaxCollatzSequenceV1
--- PASS: TestMaxCollatzSequenceV1 (0.23s)
=== RUN   TestMaxCollatzSequenceV2
--- PASS: TestMaxCollatzSequenceV2 (0.02s)
=== RUN   TestMaxCollatzSequenceV3
--- PASS: TestMaxCollatzSequenceV3 (0.03s)
=== RUN   TestMaxCollatzSequenceV4
--- PASS: TestMaxCollatzSequenceV4 (0.02s)
goos: darwin
goarch: amd64
pkg: github.com/andreasatle/ProjectEuler100/p014
BenchmarkMaxCollatzSequenceV1
BenchmarkMaxCollatzSequenceV1-12    	       5	 221600207 ns/op
BenchmarkMaxCollatzSequenceV2
BenchmarkMaxCollatzSequenceV2-12    	      82	  17817839 ns/op
BenchmarkMaxCollatzSequenceV3
BenchmarkMaxCollatzSequenceV3-12    	      33	  35346233 ns/op
BenchmarkMaxCollatzSequenceV4
BenchmarkMaxCollatzSequenceV4-12    	      67	  15559746 ns/op
PASS
ok  	github.com/andreasatle/ProjectEuler100/p014	6.270s
```

V3 gives a real improvement, but V4 is not improving much from V3 (even though its the fastest implementation.)