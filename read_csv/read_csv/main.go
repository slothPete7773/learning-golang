package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("sample.csv")

	if err != nil {
		panic("Error, file not found.")
	}

	defer file.Close()

	fmt.Println("built-in `encoding/csv` does not differentiate header row from data row(s).")
	fmt.Println("It will send back header row as an Array of item as well.")
	fmt.Println("===========================================================================")

	// A: Read all records at once.

	// records := readAll(file)
	// fmt.Println(records)

	// B: Read one record.

	// record := readOne(file)
	// fmt.Println(record)

	// C: Read one record sequentially.
	fmt.Println("Read one `Read()` is similar to cursor read.")
	fmt.Println("Cursor will sequentially move as read() is executed.")
	readOneMultiple(file)

}

func readAll(file *os.File) [][]string {
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		panic("Error while reading records.")
	}

	return records
}

func readOne(file *os.File) []string {
	reader := csv.NewReader(file)

	record, err := reader.Read()

	if err != nil {
		panic("Error while reading records.")
	}

	return record
}

func readOneMultiple(file *os.File) {
	reader := csv.NewReader(file)

	record, _ := reader.Read()
	fmt.Printf("1: %v\n", record)

	record, _ = reader.Read()
	fmt.Printf("2: %v\n", record)
}
