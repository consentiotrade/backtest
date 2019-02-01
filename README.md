# web test

## what we have, what to do

We have a rudimentary web server to upload and review stocks.
We would like to:

* add clients and client portfolios: e.g. "joe has 34 shares of AAPL, 43 of MSFT ..."
* create a current-value calculator for users
* design (not necessarily build) a bulk uploader for user portfolios
* design (not necessarily build) API for portfolio upload, portfolio viewing etc.
* comment on code (your own or the existing code), defend choices/design
* what would you do next with this project?

## howto

* compile:  from this directory `go build local/gosam/cmd/serve`
* run: from this directory: `./serve`
* open browser at [root](http://localhost:8080/)
* run tests: `go test -v local/gosam/stock/...`