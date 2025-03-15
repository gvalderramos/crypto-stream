package main

import (
	"fmt"

	"github.com/gvalderramos/crypto-stream/internal/crypto_stream_api"
)

func main() {
	fmt.Println("Starting Binance Watcher...")

	// Example usage of the internal package
	api := crypto_stream_api.NewAPI()
	data := api.FetchPrices("binance")
	fmt.Println(data)
}
