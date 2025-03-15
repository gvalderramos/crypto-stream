package crypto_stream_api

import "fmt"

// API struct to interact with different exchanges
type API struct{}

// NewAPI creates a new API instance
func NewAPI() *API {
	return &API{}
}

// FetchPrices fetches price data from the given exchange
func (api *API) FetchPrices(exchange string) string {
	return fmt.Sprintf("Fetching prices from %s...", exchange)
}
