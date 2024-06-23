package main

import "fmt"

func main() {
	fmt.Print("If-else condition statement type A:")
	var score int = 62 // ตัวอย่างสมมุติว่า 62 คะแนน

	if score >= 70 {
		fmt.Printf("PASSED")
	} else {
		// ก็จะทำงานตรงตำแหน่งนี้ เนื่องจากเงื่อนไขอันบนเป็นเท็จ
		fmt.Printf("FAILED")
	}

	fmt.Print("If-else condition statement type B:")
	// แบบที่ 2
	age := 24
	if age < 20 {
		// เมื่อเป็นจริงจะทำตรงนี้
		fmt.Printf("Alcohol ineligible.")
	} else if (age >= 20) && (age <= 45) {
		fmt.Printf("Alcohol eligible.")
		// เมื่อ condition statement อันบนเป็นเทจ และ condition statement อันนี้เป็นจริงจะทำตรงนี้
	} else {
		fmt.Printf("Alcohol should be avoided.")
		// เมื่อไม่มี condition statement ไหนเป็นจริงเลย = จะมาทำตรงนี้
	}

	fmt.Println("Switch case condition:")

	var dayOfWeek = 3

	switch dayOfWeek {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid Day")
	}

	fmt.Println("Preprocess before If-Else:")

	num1 := 5
	num2 := 10

	// Without Preprocess.
	sumNum := num1 + num2
	if sumNum >= 10 {
		fmt.Println("sumNum more than 10")
	}

	// With Preprocess (Inline).
	if sumNum := num1 + num2; sumNum >= 10 {
		fmt.Println("sumNum more than 10")
	}

}
