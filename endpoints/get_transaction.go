package endpoints

import (
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"go_test_PuTTY/utils"
	"io/ioutil"
	"log"
	"net/http"
)

// This will conform our request structure.
type ethereum_request struct {

	// The field that will be accepted on the left, and the right is the response formatted fields.
	Ethereum_ID string `json:"ethereum_ID"`
}

// This holds our definition of a list that can hold multiple ethereum request.
type eth_req_list []ethereum_request


// ENDPOINT: ("/get_transaction")
// ACCEPTS: GET
func GetTransaction(responseObject http.ResponseWriter, requestObject *http.Request) {

	// First connect
	utils.ConnectClient()

	var ethereum_ID ethereum_request

	// This is where we read the request, and deserialize into a 'slice'
	// This makes it human-readable.
	requestBody, err := ioutil.ReadAll(requestObject.Body)

	if err != nil {
		utils.LogError.Println("ERROR! You're not adhering to the schema. Can you read bro? ERROR: " + err.Error())
	}
	// Unpack the slice into a string.
	json.Unmarshal(requestBody, &ethereum_ID)

	// Set our account ID.
	account := common.HexToAddress(ethereum_ID.Ethereum_ID)

	enterTransactionID(account)
	utils.LogInfo.Printf("The received ethereum request ID is: [ %s ]", ethereum_ID)
	responseObject.WriteHeader(http.StatusCreated)

	// Finally, write the new payload to the response object.
	json.NewEncoder(responseObject).Encode(ethereum_ID)
}

// Helper method for sourcing an ethereum transaction using an ID.
func enterTransactionID(account common.Address) {

	client, err := utils.ConnectClient()
	if err != nil {
		utils.LogError.Println("Can't connect: " + err.Error())
	}
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	utils.LogInfo.Printf("The balance of the ethereum address is: %s ether", balance)
}