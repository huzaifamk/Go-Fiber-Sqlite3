package main

import (
	db "go-fiber-crud/database"
	router "go-fiber-crud/routes"

	"github.com/gofiber/fiber/v2"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {

	// Using Fiber as the web framework to start the server and handle the routes
	server := fiber.New()

	// Initialize the database connection
	db.IniDatabase()

	// Add the routes to the server
	router.AddRoutes(server)

	// Start the server on port 3000
	server.Listen(":3000")

	// Close the database connection when the server is stopped
	defer db.DBConn.Close()
}
