package main

import "fmt"

func main() {
	// Array 1
	var arrayA [3]int
	arrayA[0] = 1
	arrayA[1] = 3
	arrayA[2] = 7

	fmt.Println(arrayA)

	// Array 2
	arrayB := [3]int{1, 2, 3}
	fmt.Println(arrayB)

	// Array 3
	arrayC := []int{12, 23, 34}
	fmt.Println(arrayC)

	// Array 4
	arrayD := [5]int{123, 441, 41}
	arrayD[3] = 99
	arrayD[4] = 1312
	fmt.Println(arrayD)
	// Go immediately notify the error out of bound.
	// arrayD[5] = 13122

	// Array 5
	arrayE := [5]int{123, 441, 41}
	fmt.Printf("Length of array: %d\n", len(arrayE))

}
