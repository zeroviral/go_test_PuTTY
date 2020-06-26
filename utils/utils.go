package utils

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)
var singletonInstance *ethclient.Client

func ConnectClient() (*ethclient.Client, error) {

	url := "https://mainnet.infura.io/v3/b1e85ed87262471ea8ce2fd3e6d2f7c8"

	if singletonInstance == nil {

		client, err := ethclient.Dial(url)
		if err != nil {
			log.Fatalf("Can't connect")
		}

		singletonInstance = client
		fmt.Println("Successfully connected to Infura/Ethereum")
	}

	return singletonInstance, nil
}
