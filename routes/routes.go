package routes

import "github.com/gofiber/fiber/v2"

const (
	GetUsersRoute = "/users"
)

func AddRoutes(s *fiber.App) {
	s.Get(GetUsersRoute, GetUsers)
}