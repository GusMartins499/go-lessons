// GOROUTINES
//
// A goroutine is a function running concurrently with the rest of the
// program. Start one by putting "go" before a call: go doSomething().
//
// Rules:
//   - They are very lightweight (a few KB of stack), managed by the Go
//     runtime, not by the OS. Running thousands of them is normal.
//   - "main" is also a goroutine. When main returns, the program ENDS,
//     even if other goroutines are still running (they are killed).
//   - The execution order between goroutines is NOT guaranteed.
//   - The "go" statement doesn't wait and discards the return value.
//
// Here time.Sleep is used to wait, just for teaching. In real code use
// sync.WaitGroup (see the waitgroup folder) or channels.
package main

import (
	"fmt"
	"time"
)

func say(text string, times int) {
	for i := 1; i <= times; i++ {
		fmt.Printf("%s %d\n", text, i)
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	// Sequential: "first" only starts after "zero" finishes
	say("zero", 2)

	// Concurrent: both run at the same time, the output interleaves
	go say("first", 3)
	go say("second", 3)

	// Anonymous function as a goroutine
	go func() {
		fmt.Println("anonymous goroutine")
	}()

	// Without this, main would end before the goroutines print anything
	time.Sleep(100 * time.Millisecond)

	// Goroutine started right before main ends: nothing guarantees it runs.
	// Run a few times: sometimes it prints, sometimes it doesn't.
	go fmt.Println("this may or may not appear")
	fmt.Println("main finished")
}
