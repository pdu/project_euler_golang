package main

// sln1 is the most straight forward solution
//
// for loop the range of [1, 1000), sum the value if it can be divided by 3 or 5
//
// from the following logic, you can learn the usage of for loop and if condition
func sln1() int {
	sum := 0
	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}
