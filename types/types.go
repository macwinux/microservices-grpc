package types

type PriceResponse struct {
	Ticker string  `json:"ticker"`
	Price  float64 `json:"price"`
}

type key string

const (
	RequestID key = "requestID"
)
