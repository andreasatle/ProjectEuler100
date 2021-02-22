package p004

func LargestPalindrom(minNum, maxNum int) (int, int, int) {
	maxPal, maxA, maxB := 1, 1, 1
	for a := maxNum; a >= minNum; a-- {
		for b := a; b >= minNum && a*b > maxPal; b-- {
			prod := a * b
			if isPal := isPalindrom(prod); isPal {
				if prod > maxPal {
					maxPal, maxA, maxB = prod, a, b
				}
			}
		}
	}
	return maxPal, maxA, maxB
}

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
