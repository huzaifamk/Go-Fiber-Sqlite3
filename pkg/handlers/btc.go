package users

import (
	svc "go-fiber-api/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetBTCPrice(c *fiber.Ctx) error {

	// Get the price of BTC from CoinMarketCap API
	btcPrice := svc.GetBTCPrice(c)
	// Convert the price to a string
	btcPriceString := strconv.FormatFloat(btcPrice.Data.BTC.Quote.USD.Price, 'f', 2, 64)
	// Return the price of BTC
	return c.JSON("Bitcoin Price in USD: $" + btcPriceString)

}