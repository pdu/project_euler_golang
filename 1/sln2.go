package main

// multipleSum returns the sum of the values in the rang of [1, n] which can be divided by k
//
// the sum of: k, 2k, 3k, ... mk = k * (m * (m + 1) / 2)
//             m = n / k
func multipleSum(n, k int) int {
	if k > n {
		return 0
	}
	m := n / k
	return k * (1 + m) * m / 2
}

// sln2 uses math to solve the problem
//
// set_3 = the values which can be divided by 3
// set_5 = the values which can be divided by 5
// set_15 = the values which can be divided by 15
// the intersection of set_3 and set_5 is set_15
// so the sum is set_3 + set_5 - set_15
func sln2() int {
	sum := multipleSum(999, 3) + multipleSum(999, 5) - multipleSum(999, 15)
	return sum
}
