package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/ssaiganesh/go-jwt-authentication/database"
	"github.com/ssaiganesh/go-jwt-authentication/routes"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true, // for frontend to get cookie we send and send it back
	})) // used when frontend uses different port to backend

	routes.Setup(app)

	app.Listen(":8000")
}
