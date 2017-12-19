package main

// transfer the number into digits and compare the digits from beginning and from ending
func isPalinDrome1(n int) bool {
	var digits []int
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	for i := 0; i < len(digits)/2; i++ {
		if digits[i] != digits[len(digits)-1-i] {
			return false
		}
	}
	return true
}

func sln1() int {
	result := 0
	maxInt := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}
	// i loop from 100 to 999
	// j loop from i to 999
	// it can save the space of compute 100 * 500 and 500 * 100 twice
	for i := 100; i <= 999; i++ {
		for j := i; j <= 999; j++ {
			product := i * j
			if isPalinDrome1(product) {
				result = maxInt(result, product)
			}
		}
	}
	return result
}
