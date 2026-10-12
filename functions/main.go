package main

import "fmt"

func math(n1,n2 int) (int, int) {
	sum := n1+n2
	sub := n1-n2

	return sum, sub
	// return n1+n2, n1-n2
}

func sum(a int, b int) int {
	return a + b
}

func main() {
	soma := sum(3, 5)
	fmt.Println(soma)

	var f = func() {
		fmt.Println("f funcion")
	}

	f()

	// _, resultSub := math(10, 5)
	resultSum, resultSub := math(10, 5)
	fmt.Println(resultSum, resultSub)
}
