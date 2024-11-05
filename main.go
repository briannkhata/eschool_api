package main

import (
	db "github.com/briannkhata/eschool/config"
	"github.com/briannkhata/eschool/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Connect to the database
	db.Connect()

	// Initialize Fiber app
	app := fiber.New()

	// Setup routes
	routes.Setup(app)

	// Start the server on port 8001
	err := app.Listen(":8001")
	if err != nil {
		panic(err)
	}
}
