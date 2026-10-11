package main

import (
	"fmt"
	"module/auxiliary"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Hello, World!")
	auxiliary.Foo()
	err := checkmail.ValidateFormat("gustavosm994@gmail.com")
	if err != nil {
		fmt.Println("Invalid email format:", err)
	}
}
