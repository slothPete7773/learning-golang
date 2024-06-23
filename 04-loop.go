package main

import "fmt"

func main() {

	fmt.Println("Regular For-Loop")
	for i := 1; i < 10; i++ {
		fmt.Printf("number: %d", i)
	}

	i := 0
	for {
		fmt.Printf("number: %d\n", i)
		i = i + 1

		if i == 1000 {
			break
		}
	}
	fmt.Println("Above is Infinity For-Loop, but has Break condition to prevent infinity loop. lol")

	fmt.Println("Go does not have primitive While Loop, so an improvise approach is used.")

	fmt.Println("Do While Loop.")
	i = 1
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
