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

	// make approve transaction
	approveAmount := new(big.Int).Mul(amountTransfer, big.NewInt(2))
	approveTxData, err := daiAbi.Pack("approve", multicallAddress, approveAmount)
	if err != nil {
		log.Printf("Error packing approveTxData: %v\n", err)
		return nil, err
	}
	approveTx, err := utils.MakeTransaction(ctx, fromAddress, daiAddress, value, approveTxData)
	if err != nil {
		log.Printf("Error making transaction: %v\n", err)
		return nil, err
	}
	chainID, err := infuraClient.NetworkID(ctx)
	if err != nil {
		log.Printf("Error getting chain id: %v\n", err)
		return nil, err
	}
	signedApproveTx, err := types.SignTx(approveTx, types.NewEIP155Signer(chainID), senderPrivateKey)
	if err != nil {
		log.Printf("Error signing tx: %v\n", err)
		return nil, err
	}
	err = infuraClient.SendTransaction(ctx, signedApproveTx)
	if err != nil {
		log.Printf("Error sending transaction: %v\n", err)
		return nil, err
	}
	log.Printf("transaction sent: %v\n", signedApproveTx.Hash().Hex())

	// make transfer transaction
	transferData, err := daiAbi.Pack("transferFrom", fromAddress, toAddress, amountTransfer)
	transferData2, err := daiAbi.Pack("transferFrom", fromAddress, toAddress, amountTransfer)

	if err != nil {
		log.Printf("error packing transfer data: %v\n", err)
		return nil, err
	}
	calls := []multicallabi.Struct0{
		multicallabi.Struct0{
			Target:   daiAddress,
			CallData: transferData,
		},
		multicallabi.Struct0{
			Target:   daiAddress,
			CallData: transferData2,
		},
	}
	multicallTxData, err := multicallAbi.Pack("aggregate", calls)
	if err != nil {
		log.Printf("error packing multicall tx data: %v\n", err)
		return nil, err
	}
	multicallTx, err := utils.MakeTransaction(ctx, fromAddress, multicallAddress, value, multicallTxData)
	if err != nil {
		log.Printf("error making transaction: %v\n", err)
		return nil, err
	}
	chainID2, err := infuraClient.NetworkID(ctx)
	if err != nil {
		log.Printf("Error getting chain id: %v\n", err)
		return nil, err
	}
	signedMulticallTx, err := types.SignTx(multicallTx, types.NewEIP155Signer(chainID2), senderPrivateKey)
	if err != nil {
		log.Printf("Error signing tx: %v\n", err)
		return nil, err
	}
	err = infuraClient.SendTransaction(ctx, signedMulticallTx)
	if err != nil {
		log.Printf("error se transaction: %v\n", err)
	}
	log.Printf("transaction sent: %v\n", signedMulticallTx.Hash().Hex())
	return signedMulticallTx, nil
}
