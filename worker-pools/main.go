// WORKER POOLS
//
// A fixed number of goroutines (workers) consume tasks from the same
// channel. It limits concurrency: instead of 1 goroutine per task, N workers
// share the work.
//
// Structure:
//  1. "jobs" channel: where the tasks go in.
//  2. N workers: each one does "for job := range jobs".
//  3. "results" channel: where the answers go out.
//  4. Close jobs when there is nothing left to send -> the workers' range
//     ends.
//  5. A WaitGroup waits for the workers and then closes results.
//
// The order of the results is NOT the order of the jobs.
package main

import (
	"fmt"
	"sync"
	"time"
)

type Result struct {
	Job    int
	Worker int
	Value  int
}

func worker(id int, jobs <-chan int, results chan<- Result) {
	for job := range jobs {
		time.Sleep(20 * time.Millisecond) // simulates work
		results <- Result{Job: job, Worker: id, Value: job * job}
	}
}

func main() {
	const numWorkers = 3
	const numJobs = 9

	jobs := make(chan int, numJobs)
	results := make(chan Result, numJobs)

	var wg sync.WaitGroup
	for id := 1; id <= numWorkers; id++ {
		wg.Go(func() {
			worker(id, jobs, results)
		})
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // no more jobs: workers exit the range

	// Closes results when all workers finish
	go func() {
		wg.Wait()
		close(results)
	}()

	start := time.Now()
	for r := range results {
		fmt.Printf("worker %d: %d² = %d\n", r.Worker, r.Job, r.Value)
	}
	// 9 jobs of 20ms with 3 workers -> ~60ms instead of ~180ms
	fmt.Println("total:", time.Since(start).Round(10*time.Millisecond))
}
