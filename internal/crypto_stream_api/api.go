package crypto_stream_api

import (
	"encoding/json"
	"fmt"
)

type CryptoStreamEvent struct {
	BrokerApi    string
	Currency     string
	CurrentPrice float32
	LowPrice     float32
	HightPrice   float32
}

type Response interface {
	ToCryptoStreamEvent() *CryptoStreamEvent
}

func (c *CryptoStreamEvent) String() ([]byte, error) {
	res, err := json.Marshal(c)
	if err != nil {
		fmt.Println("Error while converting the event:", err)
		return nil, err
	}
	return res, nil
}
