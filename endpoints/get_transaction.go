package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"go_test_PuTTY/utils"
	"go_test_PuTTY/resources"
	"io/ioutil"
	"log"
	"net/http"
)


// ENDPOINT: ("/get_transaction")
// ACCEPTS: GET
func GetTransaction(responseObject http.ResponseWriter, requestObject *http.Request) {

	// First connect
	utils.ConnectClient()

	var ethereumRequest resources.EthereumRequest

	// This is where we read the request, and deserialize into a 'slice'
	// This makes it human-readable.
	requestBody, err := ioutil.ReadAll(requestObject.Body)

	if err != nil {
		utils.LogError.Println("ERROR! You're not adhering to the schema. Can you read bro? ERROR: " + err.Error())
	}
	// Unpack the slice into a string.
	json.Unmarshal(requestBody, &ethereumRequest)

	// Set our account ID.
	account := common.HexToAddress(ethereumRequest.EthereumId)

	balance := getBalanceByTransactionID(account)
	utils.LogInfo.Printf("The received ethereum request ID is: [ %s ]", ethereumRequest)
	responseObject.WriteHeader(http.StatusCreated)


	var ourResponse resources.EthereumResponse
	ourResponse.EthereumId = ethereumRequest.EthereumId
	ourResponse.Message = fmt.Sprintf("The current Ethereum balance is: %s", balance)
	// Finally, write the new payload to the response object.
	json.NewEncoder(responseObject).Encode(ourResponse)
}

// Helper method for sourcing an ethereum transaction using an ID.
func getBalanceByTransactionID(account common.Address) string {

	client, err := utils.ConnectClient()
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