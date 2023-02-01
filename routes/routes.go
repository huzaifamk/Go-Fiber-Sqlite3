package routes

import (
	handlers "go-fiber-api/pkg/handlers"

	"github.com/gofiber/fiber/v2"
)

// Set the routes for the server
const (
	GetBTCPriceRoute = "/api/btc/price"
)

// AddRoutes adds the routes to the server
func AddRoutes(s *fiber.App) {
	s.Get(GetBTCPriceRoute, handlers.GetBTCPrice)
}
