package main

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"txtest/abi/daiabi"
	"txtest/abi/multicallabi"
	"txtest/client"
	"txtest/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	SingleTransfer(context.Background(), "0x0A7bAB86900B6F0500C37cf91D576DA93098B419", "2.22")
}
func SingleTransfer(ctx context.Context, to string, amount string) (*types.Transaction, error) {
	senderPrivateKey, err := crypto.HexToECDSA(client.PRIVATE_KEY_1)
	daiAbi, _ := daiabi.DaiabiMetaData.GetAbi()
	multicallAbi, _ := multicallabi.MulticallabiMetaData.GetAbi()
	daiAddress := common.HexToAddress("0xaD6D458402F60fD3Bd25163575031ACDce07538D")
	multicallAddress := common.HexToAddress("0x53c43764255c17bd724f74c4ef150724ac50a3ed")
	infuraClient := client.GetClient()
	if err != nil {
		log.Printf("Error importing private key: %v\n", err)
		return nil, err
	}

	senderPublicKey := senderPrivateKey.Public()
	senderPublicKeyECDSA, ok := senderPublicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Printf("Error extract public key from sender private key: %v\n", senderPrivateKey)
	}
	fromAddress := crypto.PubkeyToAddress(*senderPublicKeyECDSA)

	value := big.NewInt(0)

	toAddress := common.HexToAddress(to)

	amountTransfer := utils.GetAmountFromTokenAmount(amount, 18)
	log.Printf("transfer amount: %v", amountTransfer)
	daiData, _ := daiAbi.Pack("transfer", toAddress, amountTransfer)

	calls := []multicallabi.Struct0{
		multicallabi.Struct0{
			Target:   daiAddress,
			CallData: daiData,
		},
		multicallabi.Struct0{
			Target:   daiAddress,
			CallData: daiData,
		},
	}

	multicallData, err := multicallAbi.Pack("aggregate", calls)
	if err != nil {
		log.Printf("error packing call data: %v\n", err)
		return nil, err
	}

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
	// 	To:   &multicallAddress,
	// 	Data: multicallData,
	// })
	// if err != nil {
	// 	log.Printf("Error getting gas limit: %v\n", err)
	// 	return nil, err
	// }

	tx := types.NewTransaction(nonce, multicallAddress, value, 100000, gasPrice, multicallData)
	chainID, err := infuraClient.NetworkID(ctx)
	if err != nil {
		log.Printf("Error getting chain id: %v\n", err)
		return nil, err
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), senderPrivateKey)
	if err != nil {
		log.Printf("Error signing tx: %v\n", err)
		return nil, err
	}
	err = infuraClient.SendTransaction(ctx, signedTx)
	if err != nil {
		log.Printf("Error sending transaction: %v\n", err)
		return nil, err
	}

	log.Printf("transaction sent: %v\n", signedTx.Hash().Hex())
	return signedTx, nil
}
