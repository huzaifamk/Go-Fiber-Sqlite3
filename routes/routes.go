package routes

import (
	handlers "go-fiber-crud/pkg/users/handlers"

	"github.com/gofiber/fiber/v2"
)

const (
	GetUsersRoute   = "/api/v1/users"
	GetUserRoute    = "/api/v1/users/:id"
	CreateUserRoute = "/api/v1/users"
	UpdateUserRoute = "/api/v1/users/:id"
	DeleteUserRoute = "/api/v1/users/:id"
)

func AddRoutes(s *fiber.App) {
	s.Get(GetUsersRoute, handlers.GetUsers)
	s.Get(GetUserRoute, handlers.GetUser)
	s.Post(CreateUserRoute, handlers.CreateUser)
	s.Put(UpdateUserRoute, handlers.UpdateUser)
	s.Delete(DeleteUserRoute, handlers.DeleteUser)
}
