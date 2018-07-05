package main

import (
	"fmt"

	"github.com/stellar/go/clients/horizon"
	"github.com/stellar/go/build"
)

func main() {
	payer :="GAQPUXEXZGPSELVN4VOXBBM4HQJWG7KB3CTKAF56MMMZFIMFRIAM3BUG"
	//payerSeed :="SCOCFMCNYYJJNTTENCU452RXHB4YWXESELFKZHKWEDTVO45O4Q4CDFIK"
	from :="GA63YLH3OA4UDMB2N77FN5HJAQ2AS43BE3INQJRMFFQISBWSDN43SEWH"
	fromSeed:="SARLW5R27ZR4VTSHPEGMPTXGDC2AH7PUJQHEVMUJOAIFLP5E4BASTCZG"

	seq, err := horizon.DefaultTestNetClient.SequenceForAccount(from)
	if err != nil {
		panic(err)
	}
	fmt.Printf("seq: %d\n", seq)

	tx, err := build.Transaction(
		build.SourceAccount{from},
		build.Sequence{uint64(seq+1)},
		build.TestNetwork,
		build.SetOptions(
			//build.InflationDest(),
			//build.SetAuthRequired(),
			//build.SetAuthRevocable(),
			//build.SetAuthImmutable(),
			//build.ClearAuthRequired(),
			//build.ClearAuthRevocable(),
			//build.ClearAuthImmutable(),
			//build.MasterWeight(1),
			//build.SetThresholds(2, 3, 4),
			//build.HomeDomain("huobi"),
			build.AddSigner(payer, 6),
		),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("tx:", tx)

	txe, err := tx.Sign(fromSeed)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("txe:", txe)
	txeB64, err := txe.Base64()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("txeB64:", txeB64)

	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(txeB64)
	if err != nil {
		fmt.Println("err", err)

		fmt.Println(resp)
		fmt.Println(resp.Result)
		return
	}
	fmt.Println("Successful Transaction:")
	fmt.Println("resp.Ledger:", resp.Ledger)
	fmt.Println("resp.Hash:", resp.Hash)

}