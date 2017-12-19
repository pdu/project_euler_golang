package main

func isPalinDrome2(digits []int, n int) bool {
	m := 0
	for n > 0 {
		digits[m] = n % 10
		m++
		n /= 10
	}
	for i := 0; i < m/2; i++ {
		if digits[i] != digits[m-1-i] {
			return false
		}
	}
	return true
}

func sln2() int {
	result := 0
	maxInt := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}
	// allocate the digits array outside of the isPalinDrome function
	// it saves a lot of memory allocation time and GC time
	// due to the time, it has more than 12x performance boost
	digits := make([]int, 6)
	for i := 100; i <= 999; i++ {
		for j := i; j <= 999; j++ {
			product := i * j
			if isPalinDrome2(digits, product) {
				result = maxInt(result, product)
			}
		}
	}
	return result
}
