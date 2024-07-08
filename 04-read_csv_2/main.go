package main

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

type Book struct {
	Id     int    `csv:"number"`
	Author string `csv:"author"`
	Title  string `csv:"title"`
}

func main() {
	file, err := os.Open("sample.csv")

	if err != nil {
		panic("Error, file not found.")
	}

	defer file.Close()

	books := []*Book{}

	if err := gocsv.UnmarshalFile(file, &books); err != nil {
		panic(err)
	}

	for _, book := range books {
		fmt.Printf("Book ID: %d > \"%s\" - %s\n", book.Id, book.Title, book.Author)
	}

	fmt.Println("================ End ================")
	fmt.Println("================ Start Stringify object to CSV string. ================")

	if _, err := file.Seek(0, 0); err != nil {
		panic(err)
	}

	books = append(books, &Book{Id: 12, Title: "Another new book Vol. 12", Author: "Martin"}) // Add books
	books = append(books, &Book{Id: 13, Title: "Another new book Vol. 13"})
	books = append(books, &Book{Id: 14, Title: "Another new book Vol. 14", Author: "Martin"})
	books = append(books, &Book{Id: 15, Title: "Another new book Vol. 15"})
	csvContent, err := gocsv.MarshalString(&books) // Get all clients as CSV string
	//err = gocsv.MarshalFile(&clients, clientsFile) // Use this to save the CSV back to the file
	if err != nil {
		panic(err)
	}
	fmt.Println(csvContent) // Display all clients as CSV string
}
