package main

import (
	"fmt"
	"log"

	"example.com/slothpete7773/gorm-tutorial/book"
	"example.com/slothpete7773/gorm-tutorial/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("books.db"))

	if err != nil {
		panic("Failed to connect to database.")
	}
	fmt.Println("Connection Opened to Database")
	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	setupDatabase()

	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)

	log.Fatal(app.Listen("localhost:8000"))
}
