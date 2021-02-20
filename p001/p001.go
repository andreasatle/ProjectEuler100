package p001

import "sync"

func CountV1(a, b, n int) int {
	sum := 0

	for i := 1; i < n; i++ {
		if i%a*i%b == 0 {
			sum += i
		}
	}
	return sum
}

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
