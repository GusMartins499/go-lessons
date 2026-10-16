// BUFFERED CHANNELS
//
// make(chan T, n) creates a channel with space for n values (a queue).
//   - Sending only blocks when the buffer is FULL.
//   - Receiving only blocks when the buffer is EMPTY.
//   - len(ch): how many values are waiting. cap(ch): buffer size.
//
// So it works even in a single goroutine, as long as you don't exceed the
// capacity. Sending a 3rd value into a buffer of 2 with nobody receiving
// causes a deadlock.
//
// When to use: to decouple producer and consumer (absorb bursts), or to
// limit concurrency (semaphore). Don't add a buffer just to "fix" a deadlock
// without understanding it.
package main

import (
	"fmt"
	"time"
)

func main() {
	// Works in a single goroutine: there is space in the buffer
	ch := make(chan string, 2)
	ch <- "first"
	ch <- "second"
	// ch <- "third" // deadlock: buffer full and nobody receiving
	fmt.Println("len:", len(ch), "cap:", cap(ch))
	fmt.Println(<-ch) // FIFO: first
	fmt.Println(<-ch) // second

	// Fast producer, slow consumer: the producer only waits when full
	jobs := make(chan int, 3)
	go func() {
		for i := 1; i <= 5; i++ {
			jobs <- i
			fmt.Println("sent", i)
		}
		close(jobs)
	}()
	for j := range jobs {
		time.Sleep(20 * time.Millisecond)
		fmt.Println("processed", j)
	}

	// Semaphore: at most 2 tasks at the same time
	sem := make(chan struct{}, 2)
	done := make(chan bool)
	for i := 1; i <= 4; i++ {
		go func() {
			sem <- struct{}{}        // takes a slot (blocks if 2 are taken)
			defer func() { <-sem }() // frees the slot
			fmt.Println("task", i, "running")
			time.Sleep(20 * time.Millisecond)
			done <- true
		}()
	}
	for range 4 {
		<-done
	}
}
