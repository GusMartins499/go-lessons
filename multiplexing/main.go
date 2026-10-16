// MULTIPLEXING (FAN-IN)
//
// Joins several channels into a single one. Whoever reads doesn't need to
// know how many sources exist: it just reads from the merged channel.
//
// Two ways:
//   - select: good for a FIXED and small number of channels.
//   - one goroutine per channel + WaitGroup: works for ANY number of
//     channels. The WaitGroup closes the output when all inputs are closed.
//
// The opposite is FAN-OUT: several goroutines reading from the same channel
// (that's what the worker pool does).
package main

import (
	"fmt"
	"sync"
	"time"
)

func source(name string, total int, interval time.Duration) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for i := 1; i <= total; i++ {
			time.Sleep(interval)
			ch <- fmt.Sprintf("%s %d", name, i)
		}
	}()
	return ch
}

// Fan-in with select: two fixed channels
func mergeTwo(a, b <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for a != nil || b != nil {
			select {
			case v, ok := <-a:
				if !ok {
					a = nil // closed: turns off this case
					continue
				}
				out <- v
			case v, ok := <-b:
				if !ok {
					b = nil
					continue
				}
				out <- v
			}
		}
	}()
	return out
}

// Fan-in with WaitGroup: any number of channels
func merge(channels ...<-chan string) <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup
	for _, ch := range channels {
		wg.Go(func() {
			for v := range ch {
				out <- v
			}
		})
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	fmt.Println("--- select ---")
	for msg := range mergeTwo(
		source("email", 2, 15*time.Millisecond),
		source("sms", 3, 10*time.Millisecond),
	) {
		fmt.Println(msg)
	}

	fmt.Println("--- WaitGroup ---")
	for msg := range merge(
		source("A", 2, 10*time.Millisecond),
		source("B", 2, 15*time.Millisecond),
		source("C", 2, 20*time.Millisecond),
	) {
		fmt.Println(msg)
	}
}
