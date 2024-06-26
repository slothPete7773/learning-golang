package main

import (
	"fmt"
	"log"

	"example.com/slothpete7773/gorm-tutorial/database"
	"example.com/slothpete7773/gorm-tutorial/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("models.db"))

	if err != nil {
		panic("Failed to connect to database.")
	}
	fmt.Println("Connection Opened to Database")
	database.DBConn.AutoMigrate(&model.Book{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	setupDatabase()

	app.Get("/api/v1/model", model.GetBooks)
	app.Get("/api/v1/model/:id", model.GetBook)
	app.Post("/api/v1/model", model.NewBook)
	app.Delete("/api/v1/model/:id", model.DeleteBook)

	log.Fatal(app.Listen("localhost:8000"))
}
