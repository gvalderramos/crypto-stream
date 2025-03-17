package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gvalderramos/crypto-stream/internal/crypto_stream_api"
)

type CryptoData struct {
	Ethereum struct {
		USD float32 `json:"usd"`
	} `json:"ethereum"`
}

func getCurrentEthereumValue() *crypto_stream_api.CryptoStreamEvent {
	apiUrl := &url.URL{
		Scheme:   "https",
		Host:     "api.coingecko.com",
		Path:     "/api/v3/simple/price",
		RawQuery: "ids=ethereum&vs_currencies=usd",
	}
	fmt.Println(apiUrl.String())
	req, err := http.NewRequest("GET", apiUrl.String(), nil)
	crypto_stream_api.FailOnError(err, "Unable to connect to Coingecko")

	pass := os.Getenv("COINGECKO_API_KEY")
	req.Header.Add("accept", "application/json")
	req.Header.Add("x-cg-api-key", pass)

	res, err := http.DefaultClient.Do(req)
	crypto_stream_api.FailOnError(err, "Unable to GET from Coingecko")

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	crypto_stream_api.FailOnError(err, "Unable to decode from coingecko")

	var data CryptoData
	err = json.Unmarshal(body, &data)
	crypto_stream_api.FailOnError(err, "Unable to decode the coingecko message")
	return &crypto_stream_api.CryptoStreamEvent{
		BrokerApi:    "coingecko",
		Currency:     "usd",
		CurrentPrice: data.Ethereum.USD,
	}
}

func main() {
	var forever chan struct{}
	go func() {
		for {
			message := getCurrentEthereumValue()
			crypto_stream_api.PostNewEvent(message)
			time.Sleep(3 * time.Second)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
