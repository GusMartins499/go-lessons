// GENERATOR
//
// A function that returns a receive-only channel (<-chan T) and produces the
// values in its own goroutine. Whoever calls it just reads with range.
//
// Rules:
//   - The generator owns the channel: it creates, sends and CLOSES it.
//   - Returning <-chan T stops the caller from sending or closing.
//   - Infinite generators need a way to stop (a "done" channel or context),
//     otherwise the goroutine stays blocked forever (goroutine leak).
//
// Generators can be chained into a pipeline: each stage receives a channel
// and returns another one.
package main

import "fmt"

// Finite: closes the channel at the end
func count(limit int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 1; i <= limit; i++ {
			ch <- i
		}
	}()
	return ch
}

// Infinite: stops when "done" is closed
func fibonacci(done <-chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		a, b := 0, 1
		for {
			select {
			case ch <- a:
				a, b = b, a+b
			case <-done:
				return
			}
		}
	}()
	return ch
}

// Pipeline stage: receives a channel, returns another
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

func main() {
	for n := range count(3) {
		fmt.Println("count:", n)
	}

	done := make(chan struct{})
	fib := fibonacci(done)
	for range 8 {
		fmt.Print(<-fib, " ")
	}
	fmt.Println()
	close(done) // stops the generator's goroutine

	// Pipeline: count -> square
	for n := range square(count(4)) {
		fmt.Println("square:", n)
	}
}
