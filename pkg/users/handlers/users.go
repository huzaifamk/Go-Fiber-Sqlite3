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

func GetUser(c *fiber.Ctx) error {
	
	id := c.Params("id")
	db := database.DBConn
	var user models.User
	db.Find(&user, id)
	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	
	db := database.DBConn
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return err
	}
	db.Create(&user)
	return c.JSON(user)
}
