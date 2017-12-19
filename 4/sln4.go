package main

func sln4() int {
	result := 0
	maxInt := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}
	// loop from 999 to 100
	// save the search space by break the loop if the result can not be larger
	for i := 999; i >= 100; i-- {
		if i*999 <= result {
			break
		}
		for j := 999; j >= i; j-- {
			product := i * j
			if product <= result {
				break
			}
			if isPalinDrome3(product) {
				result = maxInt(result, product)
			}
		}
	}
	return result
}
