package main

import (
	"fmt"
	"shortlink/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config := config.LoadConfigApp(".")
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})
	port := fmt.Sprintf(":%d", config.AppPort)
	app.Listen(port)
}
