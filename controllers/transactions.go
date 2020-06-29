package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"go_test_PuTTY/resources"
	"go_test_PuTTY/utils"
	validate "go_test_PuTTY/validators"
	"net/http"
)

// ENDPOINT: ("/transactions")
// ACCEPTS: GET_Transactions
func TransactionsGET(h http.ResponseWriter, req *http.Request) error {
	// Offload the bulk of the validation initially.
	validRequest, err := validate.TransactionsGET(req)
	if err != nil {
		return utils.NewError4XX(err, 405, err.Error())
	}
	response := getAndFillAccountInfo(validRequest)
	// Finally, write the new payload to the response object.
	if err = json.NewEncoder(h).Encode(response); err != nil {
		return fmt.Errorf("Unable to encode JSON response: %v", err)
	}
	return nil
}

// Helper method. Uses a valid TRANSACTION GET request to create and update the EthereumResponse object.
// TODO: We should probably include some error handling here and beef up the logic if they throw us errors.
func getAndFillAccountInfo(validRequest resources.TransactionsSchemaGET) resources.EthereumResponse {
	// convert to correct format
	account := common.HexToAddress(validRequest.EID)
	// get balance
	balance := utils.GetBalanceByTransactionID(account)
	utils.LogInfo.Printf("The received ethereum request ID is: [ %s ]", validRequest.EID)
	response := resources.EthereumResponse{
		EthereumId: validRequest.EID,
		Message:    fmt.Sprintf("The current Etehreum balance is %s", balance),
	}
	return response
}


// Validating all routes in this higher level function that routes to appropriate handler.
// Currently supporting --> GET_Transactions
func TransactionsHandler(h http.ResponseWriter, req *http.Request) error {
	switch method := req.Method; method {
	case "GET":
	return TransactionsGET(h,req)
	default:
		return utils.NewError4XX(nil, 405, "please use a supported method")
	}
}
