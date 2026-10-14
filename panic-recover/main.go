// PANIC AND RECOVER
//
// PANIC: interrupts the normal flow. The function stops, the defers run
// and the panic goes up the call stack. If nobody recovers it, the program
// ends with a stack trace.
//   - Caused by runtime errors (index out of range, nil pointer,
//     writing to a nil map...) or called manually: panic("msg").
//
// RECOVER: catches the panic and lets execution continue. It ONLY works inside
// a function called via DEFER.
//
// IMPORTANT: in Go, expected errors are handled by returning "error",
// NOT with panic. Panic is for truly exceptional situations/bugs
// (the equivalent of something "that can't happen"). Don't use panic/recover
// as an everyday try/catch.
package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) int {
	return a / b // panics if b == 0
}

// Converts a panic into an error using a named return + defer + recover
func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered from panic: %v", r)
		}
	}()
	result = divide(a, b)
	return
}

func validateAge(age int) {
	if age < 0 {
		panic("negative age is impossible")
	}
	fmt.Println("valid age:", age)
}

func runProtected(f func()) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic caught:", r)
		}
	}()
	f()
	fmt.Println("this line doesn't run if f() panics")
}

// Idiomatic way (no panic): return an error
func idiomaticDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	r, err := safeDivide(10, 2)
	fmt.Println("10/2:", r, err)

	r, err = safeDivide(10, 0)
	fmt.Println("10/0:", r, err)

	runProtected(func() { validateAge(-5) })

	runProtected(func() {
		var list []int
		fmt.Println(list[3]) // index out of range
	})

	if _, err := idiomaticDivide(1, 0); err != nil {
		fmt.Println("idiomatic way:", err)
	}

	fmt.Println("the program continued normally")
}
