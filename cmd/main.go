package main

import (
	db "go-fiber-crud/database"
	router "go-fiber-crud/routes"

	"github.com/gofiber/fiber/v2"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {

	server := fiber.New()

	db.IniDatabase()

	router.AddRoutes(server)

	server.Listen(":3000")

	defer db.DBConn.Close()
}
