package main

func sln1() int {
	n := 600851475143
	maxValue := 0
	// only need to loop until the value of sqrt(n)
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			maxValue = i
			// this step makes sure the following factors are prime number
			for n%i == 0 {
				n /= i
			}
		}
	}
	// make sure to check that n is 1 or not
	if n != 1 {
		maxValue = n
	}
	return maxValue
}
