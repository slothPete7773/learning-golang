package main

import "fmt"

func saySomething(text string) {
	fmt.Printf("I say: %s\n", text)
}

func addition(a int, b int) int {
	return a + b
}
func main() {
	// Simple function
	saySomething("Hello World")

	fmt.Println(addition(1, 2))
}
