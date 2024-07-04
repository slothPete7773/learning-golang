package main

import "fmt"

func main() {
	// Reference: https://engineering.thinknet.co.th/%E0%B8%9E%E0%B8%9A%E0%B8%81%E0%B8%B1%E0%B8%9A-defer-panic-recover-%E0%B8%9C%E0%B8%B9%E0%B9%89%E0%B8%97%E0%B8%B5%E0%B9%88%E0%B8%97%E0%B8%B3%E0%B9%83%E0%B8%AB%E0%B9%89-golang-%E0%B9%80%E0%B8%9B%E0%B9%87%E0%B8%99%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B8%97%E0%B8%B5%E0%B9%88%E0%B9%84%E0%B8%A1%E0%B9%88%E0%B9%80%E0%B8%AB%E0%B8%A1%E0%B8%B7%E0%B8%AD%E0%B8%99%E0%B9%83%E0%B8%84%E0%B8%A3-%E0%B8%95%E0%B8%AD%E0%B8%99%E0%B8%97%E0%B8%B5%E0%B9%88-2-2-b29da5766410
	// caseA()
	caseB()
}

func caseA() {
	panic("PANIKK")
	a := recover()
	fmt.Println(a)

}

func caseB() {

	defer func() {
		a := recover()
		fmt.Println(a)
	}()
	fmt.Println("When go is Panic, it reach to Stack to find if there is any code to run, then if there is `recovery()`.")
	fmt.Println("recovery() will take value send via `panic()` and continue execute the program.")
	fmt.Println("As seen text `PANIKK` is sent via `panic()`, thus `recover()` receive the value and printed out.")
	panic("PANIKK")
}
