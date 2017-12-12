package main

import "sync"

// sln2 demos the usage of goroutine, channel and sync.WaitGroup
func sln2() int {
	c := make(chan int)

	go func() {
		first, second := 1, 1
		for second <= 4000000 {
			first, second = second, first+second
			c <- second
		}
		close(c)
	}()

	sum := 0

	// sync.WaitGroup mainly used for the case of waiting the jobs to finish
	// wg.Add(n) means to wait n jobs
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		// defer will execute the logic before the function exit
		// wg.Done() means one job finish
		defer wg.Done()

		for {
			if n, ok := <-c; !ok {
				break
			} else if n%2 == 0 {
				sum += n
			}
		}
	}()

	// wait until the jobs finish
	wg.Wait()

	return sum
}
