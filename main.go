package main

import (
	"aibook/app"

	"github.com/fsnotify/fsnotify"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	web := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	app.NewConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {

			} else {
				// Config file was found but another error was produced
			}
		}
	})
	viper.WatchConfig()
	db := app.NewDB()
	defer db.Close()
	web.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	web.Listen(":" + viper.GetString("web.port"))
}
