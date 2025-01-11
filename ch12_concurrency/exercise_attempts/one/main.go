package main

import (
	"fmt"
	"sync"
)

func main() {
	attemptWaitGroup()
}

func attemptWaitGroup() {
	out := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				out <- j
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	var wg2 sync.WaitGroup
	wg2.Add(1)

	go func() {
		defer wg2.Done()
		for v := range out {
			fmt.Println(v)
		}
	}()

	wg2.Wait()
}
