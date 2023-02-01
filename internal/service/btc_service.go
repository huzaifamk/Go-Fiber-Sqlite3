package service

import (
	"encoding/json"
	models "go-fiber-api/pkg/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetBTCPrice(c *fiber.Ctx) models.CoinMarketCapResponse {

	// Extract headers from the incoming request
	headers := c.Request().Header
	// Initialize the API key name
	keyName := models.API_KEY_NAME
	// Get the API key value from the headers
	keyValue := headers.Peek(keyName)
	// Get the Accept value from the headers
	acceptValue := headers.Peek("Accept")
	// Set the URL for the CoinMarketCap API
	url := "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest?symbol=BTC&convert=USD"
	// Create a new request
	req, _ := http.NewRequest(c.Method(), url, nil)
	// Set the headers for the incoming request to the CoinMarketCap API
	req.Header.Set("X-CMC_Pro_API_Key", string(keyValue))
	req.Header.Set("Accept", string(acceptValue))
	// Create a new client
	client := &http.Client{}
	// Send the request to the CoinMarketCap API
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// Initialize the response model
	var BTCResponse models.CoinMarketCapResponse
	// Decode the response from the CoinMarketCap API
	err = json.NewDecoder(resp.Body).Decode(&BTCResponse)
	if err != nil {
		panic(err)
	}
	// Return the response from the CoinMarketCap API
	return BTCResponse
}
