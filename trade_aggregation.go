package main

import (
	"fmt"

	"github.com/stellar/go/clients/horizon"
)

func main() {
	fmt.Println("Test trade aggregation")

	trades, err := horizon.DefaultTestNetClient.LoadTradeAggregations(horizon.Asset{"credit_alphanum4", "CMA", "GD6OHWXPO7T46SEL2SNAFKGHLQZQ467U6MMDFKZ6BR3UXCINQ4BP3IZA"}, horizon.Asset{Type: "native"}, 300000, horizon.Limit(10), horizon.Order(horizon.OrderAsc), horizon.StartTime(1522318000000), horizon.EndTime(1522320000000))
	if err != nil {
		panic(err)
	}

	for i, rr := range trades.Embedded.Records {
		fmt.Println(i)
		fmt.Println(rr)
	}

	fmt.Println("Account testing finished")
}
