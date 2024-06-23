package main

import "fmt"

func main() {

	fmt.Println("Slice is a dynamic-size array.")
	var arraySlice []int

	fmt.Println("Regular append data.")
	arraySlice = append(arraySlice, 10)
	fmt.Println(arraySlice)

	fmt.Println("Append data from loop.")
	for i := 12; i <= 19; i++ {
		arraySlice = append(arraySlice, i)
	}
	fmt.Println(arraySlice)

	fmt.Println("Append multiple data at once.")
	arraySlice = append(arraySlice, 33, 36, 39)
	fmt.Println(arraySlice)

	fmt.Println("Convert Array to Slice.")
	myArray := [3]int{1, 2, 3}
	convertedSlice := myArray[:]
	fmt.Println(convertedSlice)
}
