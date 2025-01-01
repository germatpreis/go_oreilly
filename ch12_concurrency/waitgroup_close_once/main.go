package main

import (
	"fmt"
	"sync"
)

func processAndGather[T, R any](in <-chan T, processor func(T) R, num int) []R {
	out := make(chan R, num)
	var wg sync.WaitGroup
	wg.Add(num)

	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			for v := range in {
				out <- processor(v)
			}
		}()
	}

	// launching a monitoring go routine that waits till all
	// processing is done
	go func() {
		wg.Wait()
		// parallel writes to the channel and the
		// channel should only be closed once, otherwise
		// the go runtime will panic
		//
		// this is the way to do it
		close(out)
	}()

	var result []R

	// the `for-range` channel loop exits when out is closed and the buffer
	// is empty. finally the function returns with all the values.
	for v := range out {
		result = append(result, v)
	}
	return result
}

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			ch <- i
		}
		close(ch)
	}()
	results := processAndGather(ch, func(i int) int {
		return i * 2
	}, 3)
	fmt.Println(results)

}
