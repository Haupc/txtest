package main

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"txtest/abi"
	"txtest/client"
	"txtest/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	addrAndAmountEthTransfer = map[string]string{
		"0x627f2A6283616956C9810c66fec47108CE7A0b42": "0.09",
		"0x64A57E165139480e0E7fC9263D61f196903C2c09": "0.08",
		"0x5fe26031Afb7a43fBBD17b08b6E73cf9dB7939ed": "0.07",
	}

	addrAndAmountRopstenTransfers = map[string]string{
		"0x6B2C5508491D5d5F1D61fC782148Ef55b0b94f8A": "0.03",
		"0x0A7bAB86900B6F0500C37cf91D576DA93098B419": "0.02",
	}
)

func CheckError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	multiTransferERC20()
}

func multiTransferERC20() (*types.Transaction, error) {
	ctx := context.Background()
	senderPrivateKey, err := crypto.HexToECDSA(client.PRIVATE_KEY_1)
	daiAbi, _ := abi.DaiMetaData.GetAbi()
	daiAddress := common.HexToAddress("0xad6d458402f60fd3bd25163575031acdce07538d")
	multiSendAbi, _ := abi.MultisendMetaData.GetAbi()
	multiSendSCAddress := common.HexToAddress("0x6A9DCFFbE376D5e529fD9461760E815f87A53d5B")
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
	listAddr, listAmount, totalAmount := getAddressAndAmount(addrAndAmountRopstenTransfers)

	// make approve transaction
	approveTxData, err := daiAbi.Pack("approve", multiSendSCAddress, totalAmount)
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

	chainID2, err := infuraClient.NetworkID(ctx)
	if err != nil {
		log.Printf("Error getting chain id: %v\n", err)
		return nil, err
	}
	multiSendTxData, _ := multiSendAbi.Pack("multiERC20Transfer", daiAddress, listAddr, listAmount)
	multiSendTx, _ := utils.MakeTransaction(ctx, fromAddress, multiSendSCAddress, value, multiSendTxData)
	signedMultiSendTx, err := types.SignTx(multiSendTx, types.NewEIP155Signer(chainID2), senderPrivateKey)
	if err != nil {
		log.Printf("Error signing tx: %v\n", err)
		return nil, err
	}
	err = infuraClient.SendTransaction(ctx, signedMultiSendTx)
	if err != nil {
		log.Printf("error se transaction: %v\n", err)
	}
	log.Printf("transaction sent: %v\n", signedMultiSendTx.Hash().Hex())
	return signedMultiSendTx, nil
}

func multiTransferEth() {
	nodeClient := client.GetClient()
	multiSendSCAddress := common.HexToAddress("0x6A9DCFFbE376D5e529fD9461760E815f87A53d5B")
	multisendAbi, err := abi.MultisendMetaData.GetAbi()
	CheckError(err)
	privateKey, err := crypto.HexToECDSA(client.PRIVATE_KEY_1)
	CheckError(err)
	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKey)
	listAddr, listAmount, totalAmount := getAddressAndAmount(addrAndAmountRopstenTransfers)
	multiSendTxData, err := multisendAbi.Pack("multiTransfer", listAddr, listAmount)
	CheckError(err)

	multisendTx, err := utils.MakeTransaction(context.Background(), fromAddress, multiSendSCAddress, totalAmount, multiSendTxData)
	CheckError(err)
	chainID, err := nodeClient.NetworkID(context.Background())
	CheckError(err)
	signedTx, err := types.SignTx(multisendTx, types.NewEIP155Signer(chainID), privateKey)
	CheckError(err)
	err = nodeClient.SendTransaction(context.Background(), signedTx)
	CheckError(err)
	log.Println("tx sent:", signedTx.Hash().Hex())
}

func getAddressAndAmount(addrAndAmount map[string]string) ([]common.Address, []*big.Int, *big.Int) {
	addrs := make([]common.Address, 0)
	amounts := make([]*big.Int, 0)
	totalAmount := big.NewInt(0)
	for addrStr, amountStr := range addrAndAmount {
		addr := common.HexToAddress(addrStr)
		amount := utils.GetAmountFromTokenAmount(amountStr, 18)
		addrs = append(addrs, addr)
		amounts = append(amounts, amount)
		totalAmount.Add(totalAmount, amount)
	}
	return addrs, amounts, totalAmount
}
