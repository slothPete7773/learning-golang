package main

import "fmt"

func showFirst() {
	fmt.Println("Show first.")
}

func showSecond() {
	fmt.Println("Show second.")
}
func main() {

	showFirst()
	fmt.Println("Immediately push another showFirst() to the bottom of Stack to execute later.")
	defer showFirst()
	showSecond()
	showSecond()
}
