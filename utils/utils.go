package utils

import (
	"github.com/ethereum/go-ethereum/ethclient"
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
