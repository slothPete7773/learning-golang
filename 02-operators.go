package main

import "fmt"

func main() {

	a := 10
	b := 3

	fmt.Println("Arithmetic Operators.")
	fmt.Printf("a + b = <%d>\n", a+b)      // 13
	fmt.Printf("a - b = <%d>\n", a-b)      // 7
	fmt.Printf("a * b = <%d>\n", a*b)      // 30
	fmt.Printf("a / b = <%d>\n", a/b)      // 3
	fmt.Printf("a modulo b = <%d>\n", a%b) // 1

	fmt.Println("Relational Operators. (Equality/Inequality)")
	fmt.Println(a == b) // false
	fmt.Println(a != b) // true
	fmt.Println(a > b)  // true
	fmt.Println(a < b)  // false
	fmt.Println(a >= b) // true
	fmt.Println(a <= b) // false

	fmt.Println("Logical Operators.")
	c := true
	d := false
	fmt.Println(c && d) // false
	fmt.Println(c || d) // true
	fmt.Println(!c)     // false
}
