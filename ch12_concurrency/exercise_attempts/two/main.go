package main

import "fmt"

func main() {
	run()
}

func run() {
	numGoRoutines := 2

	channels := make([]chan int, numGoRoutines)

	for i := 0; i < numGoRoutines; i++ {
		channels[i] = make(chan int)
	}

	for i := 0; i < numGoRoutines; i++ {
		go func() {
			defer close(channels[i])
			for j := 0; j < 10; j++ {
				channels[i] <- j
			}
		}()
	}

	count := 2
	for count != 0 {
		select {
		case v, ok := <-channels[0]:
			if !ok {
				channels[0] = nil
				count--
				break
			}
			fmt.Println(v)
		case v, ok := <-channels[1]:
			if !ok {
				channels[1] = nil
				count--
				break
			}
			fmt.Println(v)
		}

	}
}
