package main

// POINTERS
//
// A pointer holds the memory ADDRESS of a variable.
//
//	&x   -> "address of x" (creates a pointer)
//	*T   -> type "pointer to T" (e.g. *int)
//	*p   -> "value pointed to by p" (dereference)
//
// What they're for:
//   - Go passes EVERYTHING by value (a copy). With a pointer, a function can
//     change the original variable.
//   - Avoiding copies of large structs.
//
// The zero value of a pointer is nil. Dereferencing nil causes a panic.
// There is no pointer arithmetic (p++ doesn't exist) — it's safer than C.
// The Garbage Collector manages memory; there is no free/delete.
//

import "fmt"

type Point struct{ X, Y int }

func main() {
	x := 10
	p := &x // p is *int

	fmt.Println("value of x:", x)
	fmt.Println("address of x:", p)
	fmt.Println("value via pointer:", *p)

	*p = 20 // changes x through the pointer
	fmt.Println("x after *p = 20:", x)

	// Two pointers to the same variable
	q := p
	*q = 30
	fmt.Println("x:", x, "| *p:", *p, "| *q:", *q)

	// new(T) allocates a T with its zero value and returns *T
	n := new(int)
	*n = 5
	fmt.Println("new(int):", *n)

	// Pointer to struct: Go dereferences automatically when accessing fields
	pt := &Point{1, 2}
	pt.X = 100 // equivalent to (*pt).X = 100
	fmt.Println("point:", *pt)

	// Nil pointer
	var empty *int
	fmt.Println("nil pointer:", empty == nil)
	// fmt.Println(*empty) // panic: invalid memory address or nil pointer dereference
}
