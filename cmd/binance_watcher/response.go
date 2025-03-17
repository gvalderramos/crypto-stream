package main

import (
	"encoding/json"
	"time"

	"github.com/gvalderramos/crypto-stream/internal/crypto_stream_api"
)

type BinanceResponse struct {
	EventType                string  `json:"e"`
	eventTime                int64   `json:"E"`
	Symbol                   string  `json:"s"`
	PriceChange              float32 `json:"p"`
	PriceChangePerc          float32 `json:"P"`
	WeightedAvg              float32 `json:"w"`
	PrevDayClPrice           float32 `json:"x"`
	CurrentPrice             float32 `json:"c"`
	LastTradeQnt             float32 `json:"Q"`
	BestBidPrice             float32 `json:"b"`
	BestBidQnt               float32 `json:"B"`
	BestAskPrice             float32 `json:"a"`
	BestAskQnt               float32 `json:"A"`
	OpenPrice                float32 `json:"o"`
	HightPrice               float32 `json:"h"`
	LowPrice                 float32 `json:"l"`
	TotalTradedBaseAssetVol  float32 `json:"v"`
	TotalTradedQuoteAssetVol float32 `json:"q"`
	openTime                 int64   `json:"O"`
	closeTime                int64   `json:"C"`
	FirstTradeId             uint    `json:"F"`
	LastTradeId              uint    `json:"L"`
	NumberTrades             uint    `json:"n"`
}

func NewBinanceResponse(data []byte) *BinanceResponse {
	var res BinanceResponse

	valid := json.Valid(data)
	if valid {
		json.Unmarshal(data, &res)
	}
	return &res
}

func (b *BinanceResponse) EventTime() time.Time {
	return time.UnixMilli(b.eventTime)
}

func (b *BinanceResponse) OpenTime() time.Time {
	return time.UnixMilli(b.openTime)
}

func (b *BinanceResponse) CloseTime() time.Time {
	return time.UnixMilli(b.closeTime)
}

func (b *BinanceResponse) ToCryptoStreamEvent() *crypto_stream_api.CryptoStreamEvent {
	return &crypto_stream_api.CryptoStreamEvent{
		BrokerApi:    "Binance",
		Currency:     "Tether",
		CurrentPrice: b.CurrentPrice,
		LowPrice:     b.LowPrice,
		HightPrice:   b.HightPrice,
	}
}
