package utils

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func ConnectClient(url string) {
	finalCLient, err := ethclient.Dial(url)
	if err != nil {
		log.Fatalf("Can't connect")
	}

	_ = finalCLient
	fmt.Println("Successfully connected to Infura/Ethereum")
}
