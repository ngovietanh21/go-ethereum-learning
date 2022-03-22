package accounts

import (
	"context"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func AccountBalances(client *ethclient.Client) {
	fmt.Println("----------------AccountBalances----------------")
	ctx := context.TODO()

	/*
		Reading the balance of an account is pretty simple,
		call the BalanceAt method of the client passing it the account address and optional block number.
		Setting nil as the block number will return the latest balance.
	*/
	account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	balance, err := client.BalanceAt(ctx, account, nil)
	if err != nil {
		fmt.Printf("BalanceAt_blockNumberNil_err: %s\n", err.Error())
	} else {
		fmt.Printf("BalanceAt_blockNumberNil: %v\n", balance)
	}

	/*
		Passing the block number let's you read the account balance at the time of that block.
		The block number must be a big.Int.
	*/
	blockNumber := big.NewInt(5532993)
	balance, err = client.BalanceAt(ctx, account, blockNumber)
	if err != nil {
		fmt.Printf("BalanceAt_blockNumberNotNil_err: %s\n", err.Error())
	} else {
		fmt.Printf("BalanceAt_blockNumberNotNil: %v\n", balance)
	}

	/*
		Numbers in ethereum are dealt using the smallest possible unit because they're fixed-point precision,
		which in the case of ETH it's wei. To read the ETH value you must do the calculation wei / 10^18.
		Because we're dealing with big numbers we'll have to import the native Go math and math/big packages.
		Here's how'd you do the conversion.
	*/
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Printf("ethValue: %v\n", ethValue)

	/*
		Sometimes you'll want to know what the pending account balance is,
		for example after submitting or waiting for a transaction to be confirmed.
	*/
	pendingBalance, err := client.PendingBalanceAt(ctx, account)
	if err != nil {
		fmt.Printf("PendingBalanceAt_err: %s\n", err.Error())
	} else {
		fmt.Printf("PendingBalanceAt: %v\n", pendingBalance)
	}
}
