package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/gvalderramos/crypto-stream/internal/crypto_stream_api"
)

func main() {
	wsUrl := &url.URL{
		Scheme: "wss",
		Host:   "stream.binance.com:9443",
		Path:   "/ws/btcusdt@ticker",
	}
	conn, err := crypto_stream_api.NewWebSocketConn(wsUrl)
	if err != nil {
		fmt.Println("Error while connecting with Binance websocket:", err)
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Fatalln("Error reading message:", err)
			break
		}
		// var res crypto_stream_api.Response
		res := NewBinanceResponse(message)
		// fmt.Println("Received:", string(message))
		fmt.Printf("Received: %#v\n, openTime: %#v\n", res, res.OpenTime().Local())

		// Post on event broker a new message
		crypto_stream_api.PostNewEvent(res.ToCryptoStreamEvent())
	}
}
