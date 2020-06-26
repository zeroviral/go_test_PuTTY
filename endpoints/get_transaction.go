package endpoints

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ENDPOINT: ("/get_transaction")
// ACCEPTS: GET
func GetTransAction(w http.ResponseWriter, r *http.Request) {
	// TODO: Fill in the code for sourcing transaction from Ethereum

	switch r.Method {

	case "GET":

		payload := r.URL.Query()
		fmt.Println(payload)

		fmt.Println(payload.Get("somefield"))
		//ourQuery := r.URL.Query()

	}
}

// This will conform our request structure.
type ethereum_request struct {

	// The field that will be accepted on the left, and the right is the response formatted fields.
	Ethereum_ID          string `json:"ethereum_ID"`
}

// This holds our definition of a list that can hold multiple ethereum request.
type eth_req_list []ethereum_request

func TestTransaction(responseObject http.ResponseWriter, requestObject *http.Request) {
	var ethereumRequest ethereum_request

	// This is where we read the request, and deserialize into a 'slice'
	// This makes it human-readable.
	theBody, err := ioutil.ReadAll(requestObject.Body)

	if err != nil {
		fmt.Fprint(responseObject, "ERROR! You're not adhering to the schema. Can you read bro?")
	}

	// Unpack the slice into a string.
	json.Unmarshal(theBody, &ethereumRequest)

	// Print the body to the log
	fmt.Println(json.Unmarshal(theBody, &ethereumRequest))

	responseObject.WriteHeader(http.StatusCreated)

	// Finally, write the new payload to the response object.
	json.NewEncoder(responseObject).Encode(ethereumRequest)
}


// TODO: Get the transaction ID into the payload and compute. Return proper balance.
func EnterTransaction(responseObject http.ResponseWriter, requestObject *http.Request) {
	var ethereumRequest ethereum_request

	theBody, err := ioutil.ReadAll(requestObject.Body)

	if err != nil {
		fmt.Fprint(responseObject, "Transaction not found")
	}

	json.Unmarshal(theBody, &ethereumRequest)
	responseObject.WriteHeader(200)

	json.NewEncoder(responseObject).Encode(ethereumRequest)

	//account := common.HexToAddress("0x742d35cc6634c0532925a3b844bc454e4438f44e")
	//balance, err := client.BalanceAt(context.Background(), account, nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("The balance of this ethereum address is: %s ether \n", balance)

}