package main

import (
	"context"
	"fmt"
)

func countTo(ctx context.Context, max int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < max; i++ {
			select {
			case <-ctx.Done():
				return
			case ch <- i:
			}
		}
	}()
	return ch
}

func main() {

	// WithCancel returns a copy of parent with a new Done channel. The returned
	// context's Done channel is closed when the returned cancel function is called
	// or when the parent context's Done channel is closed, whichever happens first.
	//
	// Canceling this context releases resources associated with it, so code should
	// call cancel as soon as the operations running in this Context complete.
	ctx, cancel := context.WithCancel(context.Background())

	// ensure that cancel is called when main ends
	// this closes the channel returned by `Done` and since a closed
	// channel always returns a value, it ensures that the goroutine
	// running `countTo` exits.
	defer cancel()

	// creates 10 times: a channel which is immediately used in a coroutine. The goroutine:
	// - cleans up after itself (closes the channel from within with a `defer`
	// - runs an endless a loop till 10
	// - uses the first branch of the `select` to check the channel returned by Done() - if it returns a value it exits
	// - tries to write to the channel in the second branch
	ch := countTo(ctx, 10)
	for i := range ch {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}
}
