package main

// sln3 demos the usage of goroutine and channel
//
// the whole logic is like:
//		goroutine1 --channel1--> goroutine2 --channel2--> goroutine3
func sln3() int {
	// c1 is a two way channel, it can be used to share int between two channels
	//
	// we can also define a buffered channel, for example: c = make(chan int, 100)
	// it means the channel's buffer size is 100
	//
	// the differences are:
	// 1, writes to a normal channel will block until someone reads from it
	// 2, writes to a buffered channel will write directly to the channel's buffer,
	//    it will only block when the channel's buffer is full
	// 3, read from the buffered channel or the normal channel is the same
	c1 := make(chan int)
	c2 := make(chan int)

	// start a goroutine, it's like a very cheap and lightweight thread
	go func() {
		for n := 1; n < 1000; n++ {
			// write to the channel, it will block until someone reads from it
			c1 <- n
		}
		// it's better to let the goroutine which writes to the channel close the channel
		//
		// but it's not always necessary to close the channel
		// it nobody is using it, golang GC will take it
		close(c1)
	}()

	go func() {
		// use for loop like this means loop forever
		for {
			// read from the channel, it will block until someone writes to it or the channel is closed
			// if the channel has value, ok is true; or if the channel closes, ok is false
			if n, ok := <-c1; !ok {
				// use break to terminate the for loop
				break
			} else if n%3 == 0 || n%5 == 0 {
				c2 <- n
			}
		}
		close(c2)
	}()

	sum := 0
	for {
		if n, ok := <-c2; !ok {
			break
		} else {
			sum += n
		}
	}

	return sum
}
