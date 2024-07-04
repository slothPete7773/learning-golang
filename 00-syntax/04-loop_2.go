package main

import "fmt"

func main() {

	fmt.Println("Go does not have primitive While Loop, so an improvise approach is used.")

	fmt.Println("Do While Loop.")
	i := 1
	for {
		fmt.Printf("number: %d\n", i)
		i++
		if i >= 10 {
			break
		}
	}

	fmt.Println("While Loop.")
	i = 1
	for i < 10 {
		fmt.Printf("number: %d\n", i)
		i++
	}

}
