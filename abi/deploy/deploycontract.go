package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"txtest/abi"
	"txtest/client"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
)

func CheckError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
func main() {
	nodeClient := client.GetClient()
	privateKey, _ := crypto.HexToECDSA(client.PRIVATE_KEY_1)
	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKey)
	nonce, err := nodeClient.PendingNonceAt(context.Background(), fromAddress)
	CheckError(err)
	gasPrice, err := nodeClient.SuggestGasPrice(context.Background())
	CheckError(err)
	// gasLimit, err := nodeClient.EstimateGas(context.Background(), )
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = 0
	auth.GasPrice = gasPrice

	// input := "1.0"
	address, tx, instance, err := abi.DeployMultisend(auth, nodeClient)
	CheckError(err)

	fmt.Println(address.Hex())   //
	fmt.Println(tx.Hash().Hex()) //

	_ = instance
}
