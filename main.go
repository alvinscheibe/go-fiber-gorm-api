package main

import (
	"github.com/alvinscheibe/go-fiber-api/database"
	"github.com/alvinscheibe/go-fiber-api/routes"
	"github.com/gofiber/fiber/v2"
	"log"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my app")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api", welcome)

	app.Get("/api/users", routes.GetUsers)
	app.Post("/api/users", routes.CreateUser)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
