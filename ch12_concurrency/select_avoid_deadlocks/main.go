package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		v := 1
		// writes to channel 1
		ch1 <- v
		// reads from channel 2
		v2 := <-ch2
		// the line below is never executed, as main exists and kills this
		// go routine. I need to 'clean up the goroutine' - see pg 299
		fmt.Println("goroutine: ", v, v2)
	}()

	v := 2
	var v2 int

	select {
	case ch2 <- v: // writes to channel 2
	case v2 = <-ch1: // reads from channel 1
	}

	fmt.Println(v, v2)
	// produces: 2 1, no deadlock!
}
