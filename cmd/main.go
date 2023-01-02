package main

import (
	"fmt"
	database "go-fiber-crud/database"
	models "go-fiber-crud/pkg/users/models"
	router "go-fiber-crud/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func iniDatabase() {

	var err error
	database.DBConn, err = gorm.Open("sqlite3", "crud-database.db")
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Successfully connected to database")
	database.DBConn.AutoMigrate(&models.User{})
	fmt.Println("Successfully migrated database")
}

func main() {

	server := fiber.New()

	iniDatabase()

	router.AddRoutes(server)

	server.Listen(":3000")

	server.Use(func(c *fiber.Ctx) error {
		fmt.Println(c.Path())
		return c.Next()
	})

	defer database.DBConn.Close()
}
