package users

import (
	"go-fiber-crud/database"
	models "go-fiber-crud/pkg/users/models"

	"github.com/gofiber/fiber/v2"
)

// GetUsers returns all the users in the database
func GetUsers(c *fiber.Ctx) error {

	// Get the database connection object
	db := database.DBConn
	// Create an array of users
	var users []models.User
	// Get all the users from the database and store it in the array
	db.Find(&users)
	// Return the array of users as JSON
	return c.JSON(users)

}

// GetUser returns a single user from the database
func GetUser(c *fiber.Ctx) error {

	// Get the id from the request parameters
	id := c.Params("id")
	// Get the database connection object
	db := database.DBConn
	// Create a user object
	var user models.User
	// Get the user from the database and store it in the user object
	db.Find(&user, id)
	// Return the user object as JSON
	return c.JSON(user)
}

// CreateUser creates a new user in the database
func CreateUser(c *fiber.Ctx) error {

	// Get the database connection object
	db := database.DBConn
	// Create a user object
	user := new(models.User)
	// Parse the request body and store it in the user object
	if err := c.BodyParser(user); err != nil {
		return err
	}
	// Save the user object in the database
	db.Create(&user)
	return c.JSON("User created successfully!")
}

// UpdateUser updates a user in the database
func UpdateUser(c *fiber.Ctx) error {

	// Get the id from the request parameters
	id := c.Params("id")
	// Get the database connection object
	db := database.DBConn
	// Create a user object
	var user models.User
	// Get the user from the database and store it in the user object
	db.First(&user, id)
	// Check if the user exists
	if user.Fname == "" {
		return c.Status(500).SendString("User not found with ID")
	}
	// Parse the request body and store it in the user object
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	// Save the user object in the database
	db.Save(&user)

	return c.JSON("User updated successfully!")
}

func DeleteUser(c *fiber.Ctx) error {

	// Get the id from the request parameters
	id := c.Params("id")
	// Get the database connection object
	db := database.DBConn
	// Create a user object
	var user models.User
	// Get the user from the database and store it in the user object
	db.First(&user, id)
	// Check if the user exists
	if user.Fname == "" {
		return c.Status(500).SendString("User not found with ID")
	}
	// Delete the user from the database
	db.Delete(&user)

	return c.JSON("User deleted successfully!")
}
