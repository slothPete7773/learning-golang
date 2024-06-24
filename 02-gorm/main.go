package main

import (
	"fmt"

	"example.com/slothpete7773/gorm-tutorial/book"
	"example.com/slothpete7773/gorm-tutorial/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("books.db"))

	if err != nil {
		panic("Failed to connect to database.")
	}
	fmt.Println("Connection Opened to Database")
	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")

}
