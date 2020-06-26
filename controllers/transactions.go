package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"go_test_PuTTY/resources"
	"go_test_PuTTY/utils"
	"go_test_PuTTY/validators"
	"log"
	"net/http"
)

// ENDPOINT: ("/get_transaction")
// ACCEPTS: GET
func GetTransaction(h http.ResponseWriter, req *http.Request) {
	// validate current request
	validRequest := validators.ValidateAddress(req)
	// convert to correct format
	account := common.HexToAddress(validRequest.EthereumId)
	// get balance
	balance := getBalanceByTransactionID(account)
	utils.LogInfo.Printf("The received ethereum request ID is: [ %s ]", validRequest)
	response := resources.EthereumResponse{
		EthereumId: validRequest.EthereumId,
		Message:    fmt.Sprintf("The current Etehreum balance is %s", balance),
	}
	// Finally, write the new payload to the response object.
	json.NewEncoder(h).Encode(response)
}

// Helper method for sourcing an ethereum transaction using an ID.
func getBalanceByTransactionID(account common.Address) string {

	client, err := utils.GetClientInstance()
	if err != nil {
		utils.LogError.Println("Can't connect: " + err.Error())
	}
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	returnString := fmt.Sprintf("%s", balance)
	return returnString
}
