package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	// Simplest route.
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")

	})

	// Route with parameter(s).
	app.Get("/param/:text", func(c *fiber.Ctx) error {
		return c.SendString("value: " + c.Params("text"))
	})

	// Route with Optional Parameter(s).
	// GET http://localhost:3000/john
	app.Get("/optional/:text?", func(c *fiber.Ctx) error {
		text := c.Params("text")
		if text != "" {
			return c.SendString("Optional value provided: " + text)
		}

		return c.SendString("Optional value did not provided, using default value: DEFAULT")
	})

	// Route with Wild-card parameter(s).
	// GET http://localhost:3000/api/user/john
	app.Get("/api/*", func(c *fiber.Ctx) error {
		return c.SendString("API path: " + c.Params("*"))
		// => API path: user/john
	})

	app.Listen("localhost:8000")
}
