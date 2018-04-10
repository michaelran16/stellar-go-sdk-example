package main

import (
	"fmt"

	"github.com/stellar/go/clients/horizon"
)

func main() {
	fmt.Println("Test trade aggregation")

	trades, err := horizon.DefaultTestNetClient.LoadTradeAggregation()
	if err != nil {
		panic(err)
	}

	fmt.Println("Result is ", trades.Embedded.Records)

	fmt.Println("Account testing finished")
}
