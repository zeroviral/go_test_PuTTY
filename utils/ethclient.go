package utils

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"sync"
	"sync/atomic"
)

var singletonInstance *ethclient.Client
var atomicFlag uint64
var lock = &sync.Mutex{}

func GetClientInstance() (*ethclient.Client, error) {
	// Guarantee uniqueness
	if atomic.LoadUint64(&atomicFlag) == 1 {
		return singletonInstance, nil
	}
	lock.Lock()
	defer lock.Unlock()
	url := "https://mainnet.infura.io/v3/b1e85ed87262471ea8ce2fd3e6d2f7c8"

	if atomicFlag == 0 {

		client, err := ethclient.Dial(url)
		if err != nil {
			LogError.Println("Can't connect")
		}

		singletonInstance = client
		LogInfo.Println("Successfully connected to Infura/Ethereum")

		// Set our atomic flag the first time
		atomic.StoreUint64(&atomicFlag, 1)
	}
	return singletonInstance, nil
}

// Helper method for sourcing an ethereum transaction using an ID.
func GetBalanceByTransactionID(account common.Address) string {
	client, err := GetClientInstance()
	if err != nil {
		LogError.Println("Can't connect: " + err.Error())
	}
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	returnString := fmt.Sprintf("%s", balance)
	return returnString
}
