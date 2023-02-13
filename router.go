package main

import "github.com/gofiber/fiber/v2"

func GenerateApp() *fiber.App {
	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	} )

	libGroup := app.Group("/library")
}