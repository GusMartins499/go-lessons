package auxiliary

import (
	"fmt"
)

// Foo prints a message from the auxiliary package.
func Foo() {
	 fmt.Println("Hello from auxiliary package!")
	 Bar()
}