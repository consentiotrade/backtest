package stock

import (
	"encoding/json"
	"fmt"
	"time"
)

// Stock is a representation of a stock such as AAPL, MSFT or such
type Stock struct {
	Symbol   string    `json:"symbol"`
	Price    float64   `json:"price"`
	Currency string    `json:"currency"`
	Time     time.Time `json:"time"`
}

// ParseStocks reads a set of stocks represented as JSON, and creates corresponding Stocks
func ParseStocks(js json.RawMessage) ([]Stock, error) {
	var stocks []Stock
	err := json.Unmarshal(js, &stocks)
	if err != nil {
		return nil, fmt.Errorf("could not parse stocks: %v", err)
	}
	return stocks, nil
}
