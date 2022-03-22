package main

import (
	"fmt"
	"go-ethereum-learning/accounts"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("we have a connection... %v\n", client)

	accounts.AccountBalances(client)
	accounts.GeneratingNewWallets()
	accounts.AddressCheck(client)
}
