package main

import (
	"fmt"
	"os"

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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8001"
	}
	err := app.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}
}
