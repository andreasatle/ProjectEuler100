## Problem 4 - Largest palindrome product

A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.

## Solution
In the first version fo the code, I converted the numbers to strings before checking for palindroms. This was modified to converting the number to be checked to a slice of digits (actually in the wrong order...) The slice can then easily be checked:
```go
func isPalindrom(n int) bool {
	// Convert the number into a slice if digits
	digits := []int{}
	for ; n > 0; n /= 10 {
		digits = append(digits, n%10)
	}

	// Check if the slice of digits is a palindrom
	lenDigits := len(digits)
	for i := 0; i < lenDigits/2; i++ {
		if digits[i] != digits[lenDigits-1-i] {
			return false
		}
	}
	return true
}
```

In order to find the largest palindrom p with two factors p=ab, with 100 ≤ a, b ≤ 999, we obtain the code:
```go
func LargestPalindrom(minNum, maxNum int) (int, int, int) {
	maxPal, maxA, maxB := 1, 1, 1
	for a := maxNum; a >= minNum; a-- {
		for b := a; b >= minNum && a*b > maxPal; b-- {
			prod := a * b
			if isPal := isPalindrom(prod); isPal {
				fmt.Println(a, b, prod, isPal)
				if prod > maxPal {
					maxPal, maxA, maxB = prod, a, b
				}
			}
		}
	}
	return maxPal, maxA, maxB
}
```

I found a way to optimize this even further. It can be shown that at least one of the factors a or b has to be divisible with 11. This was beautifully described in the forum once I solved the problem. At the end, I don't find it necessary to use this fact, once I cut the loop in b when the product a*b is smaller than the largest palindrom found so far. It can easily be shown that the palindrom 111111 = 481*231, i.e. there exist a 6 digit palindrom with factors in the correct range. This means that the palindom can be written
```
100000*n5 + 10000*n4 + 1000*n3 + 100*n2 + 10*n1 + n0, n5 == n0, n4 == n1, n3 == n2,
```
```
100001*n0 + 10010*n1 + 1100*n2 == 11*(9091*n0 + 910*n1 + 100*n2).
```
In other words the 6 digit palindrom is divisible by 11. Since we assume that a is bigger than (or equal) than b, we can't really use this interesting fact.
I think that I can't justify using this trick, since it requires to come up with that 111111 is a valid palindrome with factors in the correct range.

## Benchmark
The benchmark clocks to 681376 ns/op, which is 0.7 ms.