package main

import (
	"fmt"

	"github.com/stellar/go/clients/horizon"
	"net/http"
)

func main() {
	fmt.Println("Test trade aggregation")

	var trades horizon.TradeAggregationsPage
	var err error

	// test on testnet
	//var DefaultTestNetClient = &horizon.Client{
	//	URL:  "https://horizon-testnet.stellar.org",
	//	HTTP: http.DefaultClient,
	//}
	//
	//trades, err = DefaultTestNetClient.LoadTradeAggregations(
	//	horizon.Asset{"credit_alphanum4", "CMA", "GD6OHWXPO7T46SEL2SNAFKGHLQZQ467U6MMDFKZ6BR3UXCINQ4BP3IZA"},
	//	horizon.Asset{Type: "native"},
	//	300000,
	//	horizon.Limit(10),
	//	horizon.Order(horizon.OrderAsc),
	//	horizon.StartTime(1522318000000),
	//	horizon.EndTime(1522320000000),
	//)
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println("Start printing trade aggregations: ")
	//for i, rr := range trades.Embedded.Records {
	//	fmt.Println(i)
	//	fmt.Println(rr)
	//}
	//
	//fmt.Println("====================================")

	// test on localhost
	var LocalhostCoreClient = &horizon.Client{
		URL:  "http://localhost:8000",
		HTTP: http.DefaultClient,
	}

	trades, err = LocalhostCoreClient.LoadTradeAggregations(
		horizon.Asset{Type: "native"},
		horizon.Asset{"credit_alphanum4", "SLT", "GCKA6K5PCQ6PNF5RQBF7PQDJWRHO6UOGFMRLK3DYHDOI244V47XKQ4GP"},
		300000,
		horizon.Limit(10),
		horizon.Order(horizon.OrderAsc),
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Start printing trade aggregations: ")
	for i, rr := range trades.Embedded.Records {
		fmt.Println(i)
		fmt.Println(rr)
	}

	fmt.Println("testing finished")
}
