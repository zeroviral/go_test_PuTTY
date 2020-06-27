package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"go_test_PuTTY/resources"
	"go_test_PuTTY/utils"
	"go_test_PuTTY/validators"
	"net/http"
)

// ENDPOINT: ("/get_transaction")
// ACCEPTS: GET
func GetCurrentEthBalance(h http.ResponseWriter, req *http.Request) error {
	// validate current request
	validRequest, err := validators.ValidateRequest(req)
	if err != nil {
		return utils.NewHTTPError(nil, 405, "Invalid Request Format")
	}
	// convert to correct format
	account := common.HexToAddress(validRequest.EthereumId)
	// get balance
	balance := utils.GetBalanceByTransactionID(account)
	utils.LogInfo.Printf("The received ethereum request ID is: [ %s ]", validRequest.EthereumId)
	response := resources.EthereumResponse{
		EthereumId: validRequest.EthereumId,
		Message:    fmt.Sprintf("The current Etehreum balance is %s", balance),
	}
	// Finally, write the new payload to the response object.
	if err = json.NewEncoder(h).Encode(response); err != nil {
		return fmt.Errorf("Unable to encode JSON response: %v", err)
	}
	h.WriteHeader(200)
	return nil
}
