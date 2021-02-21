// Package p002 contains solutions to problem 2 in Project Euler.
//
// Problem 2 - Even Fibonacci numbers
//
// Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:
//
//  1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
//
// By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
package p002

func SumEvenFibonacciV1(n int) int {
	acc := 0
	for old, new := 1, 1; new <= n; old, new = new, old+new {
		if new%2 == 0 {
			acc += new
		}
	}
	return acc
}

func SumEvenFibonacciV2(n int) int {
	acc := 0
	for odd1, odd2 := 1, 1; odd1+odd2 <= n; odd1, odd2 = odd1+2*odd2, 2*odd1+3*odd2 {
		acc += odd1 + odd2
	}
	return acc
}
