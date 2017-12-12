package main

// sln1 is straight forward solution for calculating the Fibonacci sequence
func sln1() int {
	// declare the variables, equals to:
	//     var (
	// 	       sum    = 0
	// 	       first  = 1
	// 	       second = 2
	//     )
	sum, first, second := 0, 1, 2
	// we have 3 typical usages of for loop
	//
	// the first one:
	// for i:=0; i<100; i++ {
	//     ...
	// }
	//
	// the second one:
	// for i<100 {
	//     ...
	// }
	//
	// the third one:
	// for {
	//     ...
	// }
	for second <= 4000000 {
		if second%2 == 0 {
			sum += second
		}
		// this logic equals to:
		//     tmp := first
		//	   first = second
		//	   second += tmp
		first, second = second, first+second
	}
	return sum
}
