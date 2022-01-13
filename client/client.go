package client

import (
	"log"
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	_ethclient *ethclient.Client
	clientLock sync.Mutex
)

const (
	PRIVATE_KEY_1     = "73ef0bd5aadeb42768af81519e148b69f75dcdedce42e7a88524dfe7c5febe26"
	PRIVATE_KEY_2     = "d0170252c1dd03c1c2fd5bfced55e20ab1dcd2f8b9c5595271253440d9c43d15"
	PRIVATE_KEY_LOCAL = "c092a20cf88e3a826d6029112d6d59cd34d0817a6692fe9d1b8a131e9b8c4460"
)

func GetClient() *ethclient.Client {
	var err error
	if _ethclient == nil {
		clientLock.Lock()
		defer clientLock.Unlock()
		if _ethclient == nil {
			_ethclient, err = ethclient.Dial("https://ropsten.infura.io/v3/08de9f67b9a7443a84f4a30a521950c2")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	log.Println("Get client succeeded!")
	return _ethclient
}

func GetLocalClient() *ethclient.Client {
	var err error
	if _ethclient == nil {
		clientLock.Lock()
		defer clientLock.Unlock()
		if _ethclient == nil {
			_ethclient, err = ethclient.Dial("HTTP://127.0.0.1:7545")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	log.Println("Get client succeeded!")
	return _ethclient
}
