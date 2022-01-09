package service

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"sync"
	"txtest/abi/daiabi"
	"txtest/client"
	"txtest/utils"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	_transferService    *transferService
	transferServiceLock sync.Mutex
)

func GetTransferService() TransferService {
	if _transferService == nil {
		transferServiceLock.Lock()
		defer transferServiceLock.Unlock()
		if _transferService == nil {
			_transferService = &transferService{}
		}
	}
	return _transferService
}

type TransferService interface {
	SingleTransfer(ctx context.Context, to, amount string) (*types.Transaction, error)
}

type transferService struct {
}

func (s *transferService) SingleTransfer(ctx context.Context, to string, amount string) (*types.Transaction, error) {
	senderPrivateKey, err := crypto.HexToECDSA(client.PRIVATE_KEY_1)
	daiAbi, _ := daiabi.DaiabiMetaData.GetAbi()
	tokenAddress := common.HexToAddress("0xaD6D458402F60fD3Bd25163575031ACDce07538D")

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
	nonce, err := infuraClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Printf("Error getting nonce: %v\n", err)
		return nil, err
	}
	value := big.NewInt(0)
	gasPrice, err := infuraClient.SuggestGasPrice(ctx)
	if err != nil {
		log.Printf("Error getting gas price: %v\n", err)
		return nil, err
	}
	toAddress := common.HexToAddress(to)

	amountTransfer := utils.GetAmountFromTokenAmount(amount, 18)
	log.Printf("transfer amount: %v", amountTransfer)
	data, _ := daiAbi.Pack("transfer", toAddress, amountTransfer)

	gasLimit, err := infuraClient.EstimateGas(ctx, ethereum.CallMsg{
		To:   &tokenAddress,
		Data: data,
	})
	if err != nil {
		log.Printf("Error getting gas limit: %v\n", err)
		return nil, err
	}

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)
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

func (s *transferService) MultiTransfer(ctx context.Context, to string, amount string) (*types.Transaction, error) {
	senderPrivateKey, err := crypto.HexToECDSA(client.PRIVATE_KEY_1)
	daiAbi, _ := daiabi.DaiabiMetaData.GetAbi()
	tokenAddress := common.HexToAddress("0xaD6D458402F60fD3Bd25163575031ACDce07538D")

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
	nonce, err := infuraClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Printf("Error getting nonce: %v\n", err)
		return nil, err
	}
	value := big.NewInt(0)
	gasPrice, err := infuraClient.SuggestGasPrice(ctx)
	if err != nil {
		log.Printf("Error getting gas price: %v\n", err)
		return nil, err
	}
	toAddress := common.HexToAddress(to)

	amountTransfer := utils.GetAmountFromTokenAmount(amount, 18)
	log.Printf("transfer amount: %v", amountTransfer)
	data, _ := daiAbi.Pack("transfer", toAddress, amountTransfer)

	gasLimit, err := infuraClient.EstimateGas(ctx, ethereum.CallMsg{
		To:   &tokenAddress,
		Data: data,
	})
	if err != nil {
		log.Printf("Error getting gas limit: %v\n", err)
		return nil, err
	}

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)
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
