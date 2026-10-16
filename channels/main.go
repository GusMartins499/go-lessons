// CHANNELS
//
// A channel is a pipe for goroutines to send and receive values.
//   - Create:  ch := make(chan string)
//   - Send:    ch <- "value"
//   - Receive: v := <-ch
//
// UNBUFFERED channel (the default): send and receive must happen AT THE SAME
// TIME, like a handshake. Sending blocks until someone receives, and
// receiving blocks until someone sends. That's why sending and receiving in
// the same goroutine, one after the other, causes a deadlock:
//
//	ch <- "hi" // main blocks here forever
//	<-ch       // never reached
//	// fatal error: all goroutines are asleep - deadlock!
//
// Other rules:
//   - close(ch) signals there are no more values. Only the SENDER closes.
//   - Receiving from a closed channel returns the zero value immediately;
//     "v, ok := <-ch" tells you if it was closed (ok == false).
//   - "for v := range ch" receives until the channel is closed.
//   - Directions in parameters: chan<- T (send only), <-chan T (receive only).
package main

import "fmt"

// Can only send
func produce(ch chan<- int) {
	for i := 1; i <= 3; i++ {
		ch <- i
	}
	close(ch)
}

// Can only receive
func consume(ch <-chan int, done chan<- bool) {
	for v := range ch {
		fmt.Println("received:", v)
	}
	done <- true
}

func main() {
	// Basic: another goroutine sends, main receives
	message := make(chan string)
	go func() {
		message <- "Hello World"
	}()
	fmt.Println(<-message)

	// Channel as a "finished" signal + directions
	numbers := make(chan int)
	done := make(chan bool)
	go produce(numbers)
	go consume(numbers, done)
	<-done // waits for consume to finish

	// Reading from a closed channel
	ch := make(chan int)
	go func() {
		ch <- 42
		close(ch)
	}()
	v, ok := <-ch
	fmt.Println(v, ok) // 42 true
	v, ok = <-ch
	fmt.Println(v, ok) // 0 false (closed)
}
