package p014

import (
	"runtime"
	"sync"
)

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

// MaxCollatzSequenceV3 returns the number (not exceeding n) with the longest Collatz-sequence.
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
