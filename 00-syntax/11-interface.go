package main

import "fmt"

// Step 1. Create a common interface.
type FindArea interface {
	Area() float64
}

// Step 2. Declare Struct.
type Square struct {
	Side float64
}

type Triangle struct {
	Base   float64
	Height float64
}

// Step 3. Implement receiver method for the struct.

func (shape Square) Area() float64 {
	return shape.Side * shape.Side
}

func (shape Triangle) Area() float64 {
	return 0.5 * shape.Base * shape.Height
}

// Step 4. Create a function that accept Struct with defined interface.
func findArea(s FindArea) float64 {
	return s.Area()
	// fmt.Println(s.Area())
}

//	func (person AnotherPerson) getFullName() string {
//		return person.Name + " " + person.LastName
//	}
func main() {
	square := Square{Side: 4}
	triangle := Triangle{Height: 4, Base: 3}

	fmt.Printf("Area of Sqaure: %g\n", findArea(square))
	fmt.Printf("Area of Triangle: %g\n", findArea(triangle))
}
