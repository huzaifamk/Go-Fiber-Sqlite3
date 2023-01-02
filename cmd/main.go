package main

import (
	database "go-fiber-crud/database"
	router "go-fiber-crud/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func iniDatabase() {

	var err error
	database.DBConn, err = gorm.Open("sqlite3", "./db/test.db")
	if err != nil {
		panic("failed to connect database")
	}
	return

	// fmt.Println("Successfully connected to database")
	// database.DBConn.AutoMigrate(&models.User{})
	// fmt.Println("Successfully migrated database")
}

func main() {

	server := fiber.New()

	iniDatabase()

	router.AddRoutes(server)

	server.Listen(":3000")

	defer database.DBConn.Close()
}
