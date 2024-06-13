package main

import (
	"fmt"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	// Create channel to signal when any of input channels has been closed
	done := make(chan interface{})
	for _, c := range channels {
		// Iterate over each input channel.
		go func(ch <-chan interface{}) {
			// Wait input channel or done channel to be closed
			select {
			case <-ch:
				// Close done if input channel is closed
				close(done)
			case <-done:
			}
		}(c)
	}
	return done
}

func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func main() {
	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("done after %v\n", time.Since(start))
}
