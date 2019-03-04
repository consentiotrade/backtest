package stock_test

import (
	"encoding/json"
	"local/gosam/stock"
	"testing"
	"time"
)

func TestParseStocks(t *testing.T) {
	payload := `
[{
      "symbol": "AAPL",
      "price": 23.12,
      "currency": "USD",
      "time": "2019-01-31T12:12:21Z"
},
{
      "symbol": "AMZN",
      "price": 33.12,
      "currency": "USD",
      "time": "2019-01-31T14:19:01Z"      
},
{
      "symbol": "MSFT",
      "price": 12.2,
      "currency": "USD",
      "time": "2019-01-31T10:21:12Z"      
}]`
	sts, err := stock.ParseStocks(json.RawMessage(payload))
	if err != nil {
		t.Fatalf("could not parse stocks: %v", err)
	}
	if len(sts) != 3 {
		t.Errorf("wrong number of stocks parsed: %v", sts)
	}
	f, _ := time.Parse(time.RFC3339, "2019-01-31T12:12:21Z")
	if (sts[0] != stock.Stock{
		Symbol:   "AAPL",
		Price:    23.12,
		Currency: "USD",
		Time:     f,
	}) {
		t.Errorf("bad value for first stock: %#v", sts[0])
	}
}
