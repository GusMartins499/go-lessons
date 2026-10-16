// SELECT
//
// "select" waits on several channel operations at once and runs the FIRST
// one that is ready. It's like a switch, but for channels.
//
// Rules:
//   - If more than one case is ready, one is chosen AT RANDOM.
//   - If none is ready, it blocks until one is.
//   - "default" runs immediately if no case is ready (non-blocking).
//   - An empty select{} blocks forever.
//   - A nil channel is never ready: setting a channel to nil "turns off"
//     its case.
//
// Common uses: timeouts (time.After), cancellation (done channel / context),
// and non-blocking send/receive.
package main

import (
	"fmt"
	"time"
)

func main() {
	fast := make(chan string)
	slow := make(chan string)

	go func() {
		time.Sleep(10 * time.Millisecond)
		fast <- "fast"
	}()
	go func() {
		time.Sleep(30 * time.Millisecond)
		slow <- "slow"
	}()

	// Receives both, in the order they become ready
	for range 2 {
		select {
		case msg := <-fast:
			fmt.Println("received", msg)
		case msg := <-slow:
			fmt.Println("received", msg)
		}
	}

	// Timeout: gives up if the answer takes too long
	response := make(chan string)
	go func() {
		time.Sleep(100 * time.Millisecond)
		response <- "too late"
	}()
	select {
	case r := <-response:
		fmt.Println(r)
	case <-time.After(50 * time.Millisecond):
		fmt.Println("timeout!")
	}

	// Non-blocking with default
	messages := make(chan string)
	select {
	case m := <-messages:
		fmt.Println("message:", m)
	default:
		fmt.Println("no message, moving on")
	}

	// Cancellation: a worker stops when "done" is closed
	done := make(chan struct{})
	ticks := make(chan int)
	go func() {
		for i := 1; ; i++ {
			select {
			case <-done:
				fmt.Println("worker stopped")
				close(ticks)
				return
			case ticks <- i:
				time.Sleep(10 * time.Millisecond)
			}
		}
	}()
	for t := range ticks {
		fmt.Println("tick", t)
		if t == 3 {
			close(done)
		}
	}
}
