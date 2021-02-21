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
