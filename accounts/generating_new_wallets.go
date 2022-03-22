package accounts

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func GeneratingNewWallets() {
	fmt.Println("\n----------------GeneratingNewWallets----------------")

	/*
		Crypto package that provides the GenerateKey method for generating a random private key.
	*/
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		fmt.Printf("GeneratingNewWallets_GenerateKey_err: %s\n", err.Error())
	} else {
		fmt.Printf("GeneratingNewWallets_GenerateKey: %v\n", privateKey)
	}

	/*
		Convert it to bytes by importing the golang crypto/ecdsa package and using the FromECDSA method.
	*/
	privateKeyBytes := crypto.FromECDSA(privateKey)

	/*
		Convert it to a hexadecimal string by using the go-ethereum hexutil package which provides the Encode method which takes a byte slice.
		Then we strip of the 0x after it's hex encoded.

		This is the private key which is used for signing transactions and is to be treated like a password and never be shared,
		since who ever is in possesion of it will have access to all your funds.
	*/
	privateKeyString := hexutil.Encode(privateKeyBytes)
	fmt.Printf("GeneratingNewWallets_PrivateKey_FromECDSA_hexutil.Encode: %v\n", privateKeyString)
	fmt.Printf("GeneratingNewWallets_PrivateKey_FromECDSA_hexutil.Encode: %v\n", privateKeyString[2:])

	/*
		Similar to public key
	*/
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("GeneratingNewWallets_error casting public key to ECDSA")
	} else {
		fmt.Printf("GeneratingNewWallets_publicKeyECDSA: %v\n", publicKeyECDSA)
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	publicKeyString := hexutil.Encode(publicKeyBytes)
	fmt.Printf("GeneratingNewWallets_PublicKey_FromECDSA_hexutil.Encode: %v\n", publicKeyString)
	fmt.Printf("GeneratingNewWallets_PublicKey_FromECDSA_hexutil.Encode: %v\n", publicKeyString[4:])

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Printf("GeneratingNewWallets_Address: %v\n", address)

	/*
		The public address is simply the Keccak-256 hash of the public key,
		and then we take the last 40 characters (20 bytes) and prefix it with 0x.
		Here's how you can do it manually using the go-ethereum's crypto/sha3 Keccak256 functions.
	*/
	// hash := sha3.NewKeccak256()
	// hash.Write(publicKeyBytes[1:])
	// manualAddress := hexutil.Encode(hash.Sum(nil)[12:])
	// fmt.Printf("GeneratingNewWallets_manualAddress: %v\n", manualAddress)
}
