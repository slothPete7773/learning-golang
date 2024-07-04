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

	fmt.Println("For each in Golang can be done below.")
	array := []float32{1.2, 3.32, 6.9, 8.3, 9.51, 11.4412, 124.123421321}
	for index, element := range array {
		fmt.Printf("Index: [%d] Value: [%.3f]\n", index, element)
	}
}
