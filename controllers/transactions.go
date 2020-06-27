package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"go_test_PuTTY/resources"
	"go_test_PuTTY/utils"
	"go_test_PuTTY/validators"
	"net/http"
)

// ENDPOINT: ("/transactions")
// ACCEPTS: GET
func GetCurrentEthBalance(h http.ResponseWriter, req *http.Request) error {

	// First, validate that it's even a GET request.
	if !validators.ValidateRequestAction(h, req, `GET`) {
		errMsg := "could not validate request at /transactions "
		return errors.New(errMsg)
	}

	// validate current request
	validRequest, err := validators.ValidateRequest(req)
	if err != nil {
		return utils.NewError4XX(err, 405, "Invalid Request Format")
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
	return nil
}
