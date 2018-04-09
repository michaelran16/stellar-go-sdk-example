package main

import (
	"fmt"

	"github.com/stellar/go/clients/horizon"
)

func main() {
	acc1 := "GBIARRCBY5KVFQGS7YTKDJO7IWILDZUMSN5ILNPOYFKZTBJENXLFS2TP"
	acc2 := "GAJUTZBH5GTVDBVFAXEBFEHSVQ4D2C3ZXIJPR5I3GYM3ASHRJLGX3EFE"

	fmt.Println("Test account")

	account1, err := horizon.DefaultTestNetClient.LoadAccount(acc1)
	if err != nil {
		panic(err)
	}
	account2, err := horizon.DefaultTestNetClient.LoadAccount(acc2)
	if err != nil {
		panic(err)
	}

	fmt.Println("Balances is ", account1.Balances[0].Balance)
	fmt.Println("Balances is ", account2.Balances[0].Balance)
	fmt.Println("Signers is ", account1.Signers)
	fmt.Println("Signers is ", account2.Signers)
	fmt.Println("Sequence is ", account1.Sequence)
	fmt.Println("Sequence is ", account2.Sequence)

	fmt.Println("Account testing finished")

}
