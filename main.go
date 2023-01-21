package main

import (
	"github.com/alvinscheibe/go-fiber-api/database"
	"github.com/gofiber/fiber/v2"
	"log"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my app")
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Get("/api", welcome)

	log.Fatal(app.Listen(":3000"))
}
