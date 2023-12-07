package main

import (
	"aibook/app"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func main() {
	web := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	db := app.NewDB()
	defer db.Close()
	web.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	web.Listen(":3000")
}
