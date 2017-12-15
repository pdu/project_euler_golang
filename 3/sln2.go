package main

import "math"

// generate the prime list of value <= n
func getPrimes(n int) []int {
	var result []int
	// notPrime:
	// 	initial value: false, means is prime number
	//	if value updates to true, means is not prime number
	notPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		if !notPrime[i] {
			result = append(result, i)
			// mark all the multiples of i as not prime numbers
			for j := 2 * i; j <= n; j += i {
				notPrime[j] = true
			}
		}
	}
	return result
}

// check n is prime number or not
func isPrime(primeList []int, n int) bool {
	// only needs to check the value <= sqrt(n)
	for i := 0; i < len(primeList) && primeList[i]*primeList[i] <= n; i++ {
		if n%primeList[i] == 0 {
			return false
		}
	}
	return true
}

func sln2() int {
	n := 600851475143

	// only needs to generate the prime list <= sqrt(n)
	half := int(math.Sqrt(float64(n)))
	primeList := getPrimes(half)

	maxValue := 0
	// only needs to loop to the value <= sqrt(n)
	for i := 0; i < len(primeList) && primeList[i]*primeList[i] <= n; i++ {
		if n%primeList[i] == 0 {

			for n%primeList[i] == 0 {
				n /= primeList[i]
			}

			// if n is the prime number, it should be the largest prime factor
			if isPrime(primeList, n) {
				maxValue = n
				break
			} else {
				maxValue = primeList[i]
			}
		}
	}

	// make sure to check whether n is 1 or not
	if n != 1 {
		maxValue = n
	}

	return maxValue
}
