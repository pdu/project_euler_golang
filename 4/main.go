// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
//
// Find the largest palindrome made from the product of two 3-digit numbers.
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
	util.Duration(time.Now(), "solution 4", func() {
		fmt.Println("solution 4: ", sln4())
	})
}
