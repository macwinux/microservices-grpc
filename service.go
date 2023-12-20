package main

import (
	"context"
	"fmt"
)

// PriceFetcher is an interface that can fetch a price.
type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

// priceFetcher implements the PriceFetcher interface
type priceFetcher struct {
	next PriceFetcher
}

func NewPriceMockFetcher(next PriceFetcher) PriceFetcher {
	return &priceFetcher{
		next: next,
	}
}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return MockPriceFetcher(ctx, ticker)
}

var priceMocks = map[string]float64{
	"BTC": 20000.0,
	"ETH": 200.0,
	"GG":  100000.0,
}

func MockPriceFetcher(ctx context.Context, ticker string) (float64, error) {

	// mimic the HTTP roundtrip
	price, ok := priceMocks[ticker]
	if !ok {
		return price, fmt.Errorf("the given ticker (%s) is not supported", ticker)
	}
	fmt.Printf("providing the price for %s \n", ticker)
	return price, nil
}
