package users

import (
	"go-fiber-crud/database"
	models "go-fiber-crud/pkg/users/models"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {

	db := database.DBConn
	var users []models.User
	db.Find(&users)
	return c.JSON(users)

}
