package main

import (
	"fmt"

	b "github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
)

func main() {

	from := "GBIARRCBY5KVFQGS7YTKDJO7IWILDZUMSN5ILNPOYFKZTBJENXLFS2TP"
	fromsecret := "SBEIC2LB5YBET63NZZWOYC4SDKF76ZVVQJKXHSKFVVWVLQ6VKXJAD4NX"
	to := "GAJUTZBH5GTVDBVFAXEBFEHSVQ4D2C3ZXIJPR5I3GYM3ASHRJLGX3EFE"

	fmt.Println("Test send")

	tranx, err := b.Transaction(
		b.SourceAccount{AddressOrSeed: from},
		b.TestNetwork,
		b.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		b.Payment(
			b.Destination{AddressOrSeed: to},
			b.NativeAmount{Amount: "0.065"},
		),
	)

	fmt.Println("Sent transaction")

	if err != nil {
		panic(err)
	}

	txe, err := tranx.Sign(fromsecret)
	if err != nil {
		panic(err)
	}

	txeB64, err := txe.Base64()
	if err != nil {
		panic(err)
	}

	// convert txeB64 from XDR to readable transaction
	blob := txeB64
	fmt.Printf("blob base64: %s\n", txeB64)

	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(blob)
	if err != nil {
		panic(err)
	}

	fmt.Println("transaction posted in ledger:", resp.Ledger)

}
