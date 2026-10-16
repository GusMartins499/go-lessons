// WAITGROUP
//
// sync.WaitGroup waits for a group of goroutines to finish. It is a counter:
//   - Add(n): increments the counter by n (before starting the goroutine).
//   - Done(): decrements by 1 (usually "defer wg.Done()" in the goroutine).
//   - Wait(): blocks until the counter reaches zero.
//
// Since Go 1.25 there is also wg.Go(func), which does Add(1) + go + Done()
// for you.
//
// Rules:
//   - Call Add BEFORE the "go", never inside the goroutine (otherwise Wait
//     may run before Add and return too early).
//   - Pass the WaitGroup by pointer (*sync.WaitGroup), never by value: a
//     copy has its own counter.
//   - WaitGroup only waits. To share data between goroutines use a channel
//     or protect it with sync.Mutex.
package main

import (
	"fmt"
	"sync"
	"time"
)

func download(file string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(20 * time.Millisecond)
	fmt.Println("downloaded:", file)
}

func main() {
	var wg sync.WaitGroup

	files := []string{"a.txt", "b.txt", "c.txt"}
	for _, f := range files {
		wg.Add(1)
		go download(f, &wg)
	}
	wg.Wait()
	fmt.Println("all downloads finished")

	// wg.Go (Go 1.25+): no manual Add/Done
	var wg2 sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg2.Go(func() {
			fmt.Println("task", i)
		})
	}
	wg2.Wait()

	// Shared counter: WaitGroup waits, Mutex protects the data
	var (
		wg3     sync.WaitGroup
		mu      sync.Mutex
		counter int
	)
	for range 1000 {
		wg3.Go(func() {
			mu.Lock()
			counter++
			mu.Unlock()
		})
	}
	wg3.Wait()
	fmt.Println("counter:", counter) // always 1000 thanks to the mutex
}
