package main

import "fmt"

func main() {
	fmt.Println("Map is simply key-value Map.")

	// `make` เป็นคำสั่งสำหรับทำ memory allocation เพื่อให้แน่ใจว่ามี memory อยู่จริงสำหรับการใช้งานตัวแปร
	myMap := make(map[string]int)
	myMap["apple"] = 5
	myMap["banana"] = 10
	myMap["orange"] = 8
	fmt.Println(myMap)

	fmt.Println("Access an item.")
	fmt.Println("Apples:", myMap["apple"])

	fmt.Println("Delete an item from Map.")
	delete(myMap, "orange")
	fmt.Println(myMap)

	fmt.Println("Iterate over Map")
	// Iterate over the map
	// - `range` use as the iterating over the Iterable object.
	for key, val := range myMap {
		fmt.Printf("%s -> %d\n", key, val)
	}

	fmt.Println("Check if the key is existing or not.")
	val, ok := myMap["pear"]
	if ok {
		fmt.Println("Pear's value:", val)
	} else {
		fmt.Println("Pear not found inmap")
	}

}
