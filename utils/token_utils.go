package utils

import (
	"context"
	"log"
	"math/big"
	"txtest/client"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func GetAmountFromTokenAmount(tokenAmount string, decimal int32) *big.Int {
	tokenAmountFloat, ok := new(big.Float).SetString(tokenAmount)
	if !ok {
		log.Printf("Invalid amout: %v\n", tokenAmount)
		return big.NewInt(0)
	}
	mulDecimal := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimal)), nil)
	tokenAmountFloat.Mul(tokenAmountFloat, new(big.Float).SetInt(mulDecimal))
	amount := new(big.Int)
	tokenAmountFloat.Int(amount)
	return amount
}

func MakeTransaction(ctx context.Context, fromAddress, toAddress common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	infuraClient := client.GetClient()
	nonce, err := infuraClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Printf("Error getting nonce: %v\n", err)
		return nil, err
	}

	gasPrice, err := infuraClient.SuggestGasPrice(ctx)
	if err != nil {
		log.Printf("Error getting gas price: %v\n", err)
		return nil, err
	}

	// gasLimit, err := infuraClient.EstimateGas(ctx, ethereum.CallMsg{
	// 	To:    &toAddress,
	// 	Data:  data,
	// 	Value: value,
	// 	From:  fromAddress,
	// })
	// if err != nil {
	// 	log.Printf("Error getting gas limit: %v\n", err)
	// 	return nil, err
	// }
	gasLimit := uint64(3000000)
	return types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data), nil
}

func MakeLocalTransaction(ctx context.Context, fromAddress, toAddress common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	nodeClient := client.GetLocalClient()

	log.Printf("value: %v\n", value)
	nonce, err := nodeClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Printf("Error getting nonce: %v\n", err)
		return nil, err
	}
	log.Printf("none: %v\n", nonce)

	gasPrice, err := nodeClient.SuggestGasPrice(ctx)
	log.Printf("gas price: %v\n", gasPrice)
	if err != nil {
		log.Printf("Error getting gas price: %v\n", err)
		return nil, err
	}

	gasLimit, err := nodeClient.EstimateGas(ctx, ethereum.CallMsg{
		To:    &toAddress,
		From:  fromAddress,
		Data:  data,
		Value: value,
	})
	if err != nil {
		log.Printf("Error getting gas limit: %v\n", err)
		return nil, err
	}
	log.Printf("gas limit: %v\n", gasLimit)
	return types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data), nil
}
