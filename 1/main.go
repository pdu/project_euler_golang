// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
//
// Find the sum of all the multiples of 3 or 5 below 1000.
package main

import (
	"fmt"
	"time"

	"github.com/pdu/project_euler_golang/util"
)

func main() {
	util.Duration(time.Now(), "solution 1", func() {
		fmt.Println("solution 1: ", sln1())
	})

	util.Duration(time.Now(), "solution 2", func() {
		fmt.Println("solution 2: ", sln2())
	})

	util.Duration(time.Now(), "solution 3", func() {
		fmt.Println("solution 3: ", sln3())
	})
}
