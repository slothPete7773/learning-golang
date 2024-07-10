package main

import (
	"github.com/gofiber/fiber/v2"
)

func getCounter(c *fiber.Ctx) error {
	return c.Status(fiber.StatusForbidden).SendString("Some fake error")
}

func main() {
	// repo, _ =
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// app := fiber.New()

	// app.Get("/", getCounter)

	// log.Fatal(app.Listen(":8081"))
}
