package main

import (
	"log"
	"os"
)

func main() {
	oldLocation := "src/hi.txt"
	newLocation := "dest/hi.txt"
	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		log.Fatal(err)
	}
}
