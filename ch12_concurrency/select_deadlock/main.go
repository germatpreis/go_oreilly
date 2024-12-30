package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		inGoroutine := 1
		// write to channel 1
		// since this is by default an unbuffered channel,
		// THIS BLOCKS till 'main' reads from channel 1
		// (which is only done in 'main' AFTER the value 2
		// is written to channel 2)
		ch1 <- inGoroutine

		// read from channel 2
		fromMain := <-ch2

		fmt.Println("goroutine: ", inGoroutine, fromMain)
	}()

	inMain := 2
	// write into channel 2
	// this BLOCKS till the goroutine READS from channel 2 which
	// is only done after the write to channel 1 can happen
	//
	// DEADLOCK
	ch2 <- inMain

	// read from channel 1
	fromGoroutine := <-ch1

	fmt.Println("main: ", inMain, fromGoroutine)

	// produces: fatal error: all goroutines are asleep - deadlock!
}
