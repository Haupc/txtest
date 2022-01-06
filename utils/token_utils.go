package utils

import (
	"log"
	"math/big"
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
