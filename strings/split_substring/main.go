package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Hello, This is Pete, Iam sleepy."
	splited_text := strings.Split(text, ", ")

	fmt.Printf("Splited: %v\n", splited_text)
	fmt.Printf("Splited[0]: %v\n", splited_text[0])
	fmt.Printf("Splited[1]: %v\n", splited_text[1])
	fmt.Printf("Splited[2]: %v\n", splited_text[2])

	fmt.Printf("Splited[-1]: %v\n", splited_text[len(splited_text)-1])

}
