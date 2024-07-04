package main

import "fmt"

func avgArray(arr []float32) float32 {
	var sum float32 = 0
	for _, item := range arr {
		sum = sum + item
	}
	return sum / float32(len(arr))
}

func main() {
	// Simple function
	array := []float32{1.2, 3.3, 12.4, 32.9, 44.33245}

	avg := avgArray(array)

	fmt.Printf("Average of %v : %f\n", array, avg)
}
