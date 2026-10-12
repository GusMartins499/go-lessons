package main

import "fmt"

func main() {
	var variableNameString string = "Hello, World!"
	anotherVariableInt := 42

	fmt.Println(variableNameString, anotherVariableInt)

	var (
		variableString string = "This is a string variable"
		variableInt    int    = 100
		variableBool   bool   = true
	)
	fmt.Println(variableString, variableInt, variableBool)

	variable6, variable7 := "Variable 6", "Variable 7"
	fmt.Println(variable6, variable7)

	const constantVariable = "I am a constant"
	fmt.Println(constantVariable)
}
