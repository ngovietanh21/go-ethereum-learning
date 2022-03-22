package accounts

import (
	"context"
	"fmt"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func isValidEthereumAddress(address string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(address)
}

func isContract(addressString string, client *ethclient.Client) (bool, error) {
	address := common.HexToAddress(addressString)
	bytecode, err := client.CodeAt(context.Background(), address, nil)
	if err != nil {
		return false, err
	}
	isContract := len(bytecode) > 0
	return isContract, nil
}

func isEthereumAccount(addressString string, client *ethclient.Client) (bool, error) {
	isContract, err := isContract(addressString, client)
	return !isContract, err
}

func AddressCheck(client *ethclient.Client) {
	fmt.Println("\n----------------AddressCheck----------------")

	testCaseIsValidEthereumAddress := []string{
		"0x323b5d4c32345ced77393b3530b1eed0f346429d",
		"0xZYXb5d4c32345ced77393b3530b1eed0f346429d",
	}
	for i, address := range testCaseIsValidEthereumAddress {
		fmt.Printf(
			"testCaseIsValidEthereumAddress_%d_isValidEthereumAddress: %v\n",
			i, isValidEthereumAddress(address),
		)
	}

	testCaseIsContract := []string{
		"0xe41d2489571d322189246dafa5ebde1f4699f498",
	}
	for i, address := range testCaseIsContract {
		isContract, err := isContract(address, client)
		if err != nil {
			fmt.Printf("testCaseIsContract_%d_err: %v\n", i, err)
		} else {
			fmt.Printf("testCaseIsContract_%d: %v\n", i, isContract)
		}
	}

	testCaseIsEthereumAccount := []string{
		"0x8e215d06ea7ec1fdb4fc5fd21768f4b34ee92ef4",
	}
	for i, address := range testCaseIsEthereumAccount {
		isContract, err := isEthereumAccount(address, client)
		if err != nil {
			fmt.Printf("testCaseIsEthereumAccount_%d_err: %v\n", i, err)
		} else {
			fmt.Printf("testCaseIsEthereumAccount_%d: %v\n", i, isContract)
		}
	}
}
