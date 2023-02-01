package models

type CoinMarketCapResponse struct {
	Data struct {
		BTC struct {
			Quote struct {
				USD struct {
					Price float64 `json:"price"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"BTC"`
	} `json:"data"`
}

const (
	API_KEY_NAME = "X-CMC_Pro_API_Key"
)
