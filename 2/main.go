// Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:
//
//                                                 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
//
// By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
package main

import "fmt"

func main() {
	fmt.Println("solution 1: ", sln1())
	fmt.Println("solution 2: ", sln2())
}
