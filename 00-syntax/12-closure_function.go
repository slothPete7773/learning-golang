package main

import "fmt"

func main() {
	sumArray := func(arr []int) int {
		sum := 0
		for _, item := range arr {
			sum += item
		}
		return sum
	}

	array := []int{1, 2, 5, 7, 10}
	sum := sumArray(array)
	fmt.Printf("Total of array %v: %d\n", array, sum)

}
