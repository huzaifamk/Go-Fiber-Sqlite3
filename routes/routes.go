package routes

import (
	handlers "go-fiber-crud/pkg/users/handlers"

	"github.com/gofiber/fiber/v2"
)

const (
	GetUsersRoute   = "/api/v1/users"
	GetUserRoute    = "/api/v1/user/:id"
	CreateUserRoute = "/api/v1/user"
	UpdateUserRoute = "/api/v1/user/:id"
	DeleteUserRoute = "/api/v1/user/:id"
)

func AddRoutes(s *fiber.App) {
	s.Get(GetUsersRoute, handlers.GetUsers)
	s.Get(GetUserRoute, handlers.GetUser)
	s.Post(CreateUserRoute, handlers.CreateUser)
	s.Put(UpdateUserRoute, handlers.UpdateUser)
	s.Delete(DeleteUserRoute, handlers.DeleteUser)
}
