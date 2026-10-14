// DEFER
//
// "defer call()" schedules the call to run when the current function
// FINISHES (whether by a normal return or by a panic).
//
// Rules:
//   - Multiple defers run in REVERSE order (stack / LIFO).
//   - The ARGUMENTS are evaluated at the moment of the defer, not when it runs.
//   - Classic use: releasing resources right after acquiring them
//     (closing a file or connection, unlocking a mutex), guaranteeing cleanup
//     even if the function returns early. Similar to "finally" in JS.
package main

import (
	"fmt"
	"os"
	"time"
)

func order() {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")
	fmt.Println("function body")
}

func argumentsEvaluatedImmediately() {
	x := 10
	defer fmt.Println("value in defer:", x) // captures 10 now
	x = 20
	fmt.Println("current value:", x)
}

func readFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close() // will be closed on any exit from the function

	info, err := f.Stat()
	if err != nil {
		return err // f.Close() runs here too
	}
	fmt.Println("file:", info.Name(), "size:", info.Size(), "bytes")
	return nil
}

func measureTime(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func slowTask() {
	defer measureTime("slowTask")() // note the () at the end: runs measureTime now
	time.Sleep(50 * time.Millisecond)
}

func main() {
	order()
	argumentsEvaluatedImmediately()

	if err := readFile("go.mod"); err != nil {
		fmt.Println("error:", err, "(run from the defer folder)")
	}

	slowTask()
}
