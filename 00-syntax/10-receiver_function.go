package main

import "fmt"

func saySomething(text string) {
	fmt.Printf("I say: %s\n", text)
}

func addition(a int, b int) int {
	return a + b
}

type NewPerson struct {
	Name     string
	LastName string
	Age      int
}

// Receiver function
func (person NewPerson) getFullName() string {
	return person.Name + " " + person.LastName
}
func main() {
	// Simple function
	// saySomething("Hello World")

	// fmt.Println(addition(1, 2))
	fmt.Println("Receiver function is a way to bind a Method to a Struct object.")
	person := NewPerson{
		Name:     "Vee",
		LastName: "Rakit",
		Age:      12,
	}
	fmt.Println(person.getFullName())
}
