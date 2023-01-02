package routes

import (
	handlers "go-fiber-crud/pkg/users/handlers"

	"github.com/gofiber/fiber/v2"
)

const (
	GetUsersRoute = "/api/v1/users"
)

func AddRoutes(s *fiber.App) {
	s.Get(GetUsersRoute, handlers.GetUsers)
}