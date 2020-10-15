package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/consentiotrade/backtest/stock"
)

func main() {
	fmt.Println("hello world!")
	var stocks []stock.Stock

	http.HandleFunc("/stocks", func(w http.ResponseWriter, r *http.Request) {
		enc := json.NewEncoder(w)
		err := enc.Encode(stocks)
		if err != nil {
			http.Error(w, "bad stocks", http.StatusInternalServerError)
		}
	})
	http.HandleFunc("/load-stocks", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "could not parse request: "+err.Error(), http.StatusBadRequest)
			return
		}
		sts, err := stock.ParseStocks(json.RawMessage(r.Form.Get("stocks")))
		if err != nil {
			http.Error(w, "could not parse JSON: "+err.Error(), http.StatusBadRequest)
			return
		}
		stocks = sts
		w.Write([]byte("stocks stored"))
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world!"))
	})
	fmt.Println("starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
