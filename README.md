# web test

## what we have, what to do

We have a rudimentary web server to upload and review stocks.
We would like to:

* add clients and client portfolios: e.g. "Joe has: 34 shares of AAPL, 43 of MSFT ..."
* create a current-value calculator for users: e.g. "Joe's portfolio is worth 234.25 EUR"
* design (talk about, not code) a bulk uploader for user portfolios
* design (talk about, not code) API for portfolio upload, portfolio viewing etc.
* comment on code (your own or the existing code), positive and negative aspects
* defend choices/design in the previous points
* what would you do next with this project?

## howto

* if you haven't already, [install Go](https://golang.org/dl/)
* compile:  from this directory `go build local/gosam/cmd/serve`
* run: from this directory: `./serve`
* open browser at [root](http://localhost:8080/)
* run tests: `go test -v local/gosam/stock/...`