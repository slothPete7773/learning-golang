package main

import "fmt"

func main() {
	// Variable declaration type A:
	var initialVariable string
	fmt.Printf("Initial variable: \"%s\"", initialVariable)

	// Variable declaration type B:
	var nameVariable string = "slothPete"
	fmt.Printf("\nName: \"%s\"", nameVariable)

	// Variable declaration type C:
	// - Without `var`.
	// - Required `:=` instead.
	age := 20
	fmt.Printf("\nAge: %d", age)

	// Reassign variable value.
	// - Regular `=`.
	age = 24
	fmt.Printf("\nReassign age: %d\n", age)
}
