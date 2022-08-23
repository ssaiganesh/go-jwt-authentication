package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ssaiganesh/go-jwt-authentication/database"
	"github.com/ssaiganesh/go-jwt-authentication/routes"
)

func main() {
	database.Connect()

	app := fiber.New()
	routes.Setup(app)
	app.Listen(":8000")
}
