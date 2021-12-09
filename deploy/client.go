package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func client() {
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = client // we'll use this in the upcoming sections

	// ctx := context.Background()
	// tx, pending, _ := client.TransactionByHash(ctx, common.HexToHash("0x985743651d15750370335dc91e8de1bced014e0b9bd0163cd4dc7c287230438b"))
	// if !pending {
	// 	fmt.Println(tx)
	// }

}
