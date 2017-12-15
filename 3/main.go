// The prime factors of 13195 are 5, 7, 13 and 29.
//
// What is the largest prime factor of the number 600851475143 ?
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
}
