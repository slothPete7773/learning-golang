package main

import "fmt"

type Person struct {
	Name   string
	Height int
	Weight int
}

type PersonEx struct {
	Name    string
	Height  int
	Weight  int
	Address Address
}

type Address struct {
	Street  string
	City    string
	ZipCode int
}

func main() {
	// Create struct type A:
	var personA Person
	personA.Name = "slothPete"
	personA.Weight = 88
	personA.Height = 180
	fmt.Println(personA)

	// Create struct type B:
	personB := Person{
		Name:   "Veerakit",
		Weight: 99,
		Height: 199,
	}
	fmt.Println(personB)

	// Struct with Array/Slice
	fmt.Println("Struct with Array/Slice")
	var listPerson [3]Person
	listPerson[0] = Person{
		Name:   "Gaabit",
		Weight: 12,
		Height: 222,
	}
	fmt.Println(listPerson)

	// Struct with Map
	fmt.Println("Struct with Map")
	mapPerson := make(map[string]Person)

	mapPerson["p01"] = Person{Name: "Alcon", Weight: 60, Height: 180}
	fmt.Println(mapPerson["p01"])
	fmt.Println(mapPerson)

	// Nested Struct
	fmt.Println("Nested Struct")
	var person PersonEx
	person.Name = "Alice"
	person.Weight = 999
	person.Height = 999
	// person.Age = 30
	person.Address = Address{
		Street:  "123 Main St",
		City:    "Gotham",
		ZipCode: 12345,
	}

	// Alternative way to initialize a struct
	bob := PersonEx{
		Name: "Bob",
		// Age:  25,
		Weight: 123,
		Height: 9921,
		Address: Address{
			Street:  "456 Elm St",
			City:    "Metropolis",
			ZipCode: 67890,
		},
	}

	// Print struct values
	fmt.Println(person)
	fmt.Println(bob)
}
