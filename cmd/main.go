package main

import (
	router "go-fiber-api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {

	// Using Fiber as the web framework to start the server and handle the routes
	server := fiber.New()
	
	// Use the logger to log the requests
	server.Use(logger.New())

	// Add the routes to the server
	router.AddRoutes(server)

	// Start the server on port 3000
	server.Listen(":3000")
}
