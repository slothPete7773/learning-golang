package model

import (
	// "example.com/slothpete7773/gorm-tutorial/database"
	// "github.com/gofiber/fiber/v2"
	"example.com/slothpete7773/gorm-tutorial/database"
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	Title string `json:"title"`
	Name  string `json:"name"`
}

// type Company struct {
// 	ID   int
// 	Name string
// }

// type Department struct {
// 	ID   int
// 	Name string
// }

func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn
	var books []Model
	db.Find(&books)
	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book Model
	db.Find(&book, id)
	return c.JSON(book)
}

func NewBook(c *fiber.Ctx) error {
	db := database.DBConn
	book := new(Model)
	if err := c.BodyParser(book); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Create(&book)
	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var book Model
	db.First(&book, id)
	if book.Title == "" {
		return c.Status(500).SendString("No Model Found with ID")
	}
	db.Delete(&book)
	return c.SendString("Model Successfully deleted")
}
