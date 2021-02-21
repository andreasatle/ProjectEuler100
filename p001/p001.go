// Package p001 contains solutions to Problem 1 in Project Euler.
//
// Project Euler Problem 1 - Multiples of 3 and 5
//
// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
//
// Find the sum of all the multiples of 3 or 5 below 1000.
package p001

import "sync"

// CountV1 returns the sum of all the multiples of a and b below n.
// This is a non-optimized algorithm.
func CountV1(a, b, n int) int {
	sum := 0

	for i := 1; i < n; i++ {
		if i%a*i%b == 0 {
			sum += i
		}
	}
	return sum
}

// CountV2 returns the sum of all the multiples of a and b below n.
// This is an optimized algorithm.
func CountV2(a, b, n int) int {
	sum := 0

	for i := a; i < n; i += a {
		sum += i
	}

	for i := b; i < n; i += b {
		sum += i
	}

	for i := a * b; i < n; i += a * b {
		sum -= i
	}
	return sum
}

// CountV3 returns the sum of all the multiples of a and b below n.
// This is an non-optimized concurrent algorithm.
func CountV3(a, b, n int) int {
	sum := 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(3)
	go func(sum *int) {
		s := 0
		for i := a; i < n; i += a {
			s += i
		}
		mu.Lock()
		*sum += s
		mu.Unlock()
		wg.Done()
	}(&sum)

	go func(sum *int) {
		s := 0
		for i := b; i < n; i += b {
			s += i
		}
		mu.Lock()
		*sum += s
		mu.Unlock()
		wg.Done()
	}(&sum)

	go func(sum *int) {
		s := 0
		for i := a * b; i < n; i += a * b {
			s += i
		}
		mu.Lock()
		*sum -= s
		mu.Unlock()
		wg.Done()
	}(&sum)

	wg.Wait()

	return sum
}
