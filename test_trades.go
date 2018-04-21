package main

import (
	"fmt"

	"github.com/stellar/go/clients/horizon"
	"net/http"
)

func main() {
	fmt.Println("Test trades endpoint")

	var trades horizon.TradesPage
	var err error

	// test on localhost
	var LocalhostCoreClient = &horizon.Client{
		URL:  "http://localhost:8000",
		HTTP: http.DefaultClient,
	}

	trades, err = LocalhostCoreClient.LoadTrades(
		horizon.Asset{Type: "native"},
		horizon.Asset{"credit_alphanum4", "SLT", "GCKA6K5PCQ6PNF5RQBF7PQDJWRHO6UOGFMRLK3DYHDOI244V47XKQ4GP"},
		0,
		300000,
		horizon.Limit(3), horizon.Order(horizon.OrderAsc),
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Start printing trades: ")
	for i, rr := range trades.Embedded.Records {
		fmt.Println(i)
		fmt.Println(rr.Links)
		rr.Links.Self = horizon.Link{}
		rr.Links.Base = horizon.Link{}
		rr.Links.Counter = horizon.Link{}
		rr.Links.Operation = horizon.Link{}
		fmt.Println(rr)
	}

	fmt.Println("testing finished")
}
