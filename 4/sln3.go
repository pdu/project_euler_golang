package main

// revert the number and compare the value
func isPalinDrome3(n int) bool {
	m, k := 0, n
	for k > 0 {
		m = m*10 + k%10
		k /= 10
	}
	return m == n
}

func sln3() int {
	result := 0
	maxInt := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}
	for i := 100; i <= 999; i++ {
		for j := i; j <= 999; j++ {
			product := i * j
			if isPalinDrome3(product) {
				result = maxInt(result, product)
			}
		}
	}
	return result
}
