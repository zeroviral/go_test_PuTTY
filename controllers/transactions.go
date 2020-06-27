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
func GetCurrentEthBalance(h http.ResponseWriter, req *http.Request) {
	// validate current request
	validRequest, err := validators.ValidateRequest(req)
	if err != nil {
		errMsg := utils.GenerateError("Validation error has occurred", err)
		utils.LogError.Print(errMsg)
		fmt.Printf("%v", errMsg.Data)
		h.WriteHeader(422)
		json.NewEncoder(h).Encode(fmt.Sprint(errMsg))
		return
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
	newErr := json.NewEncoder(h).Encode(response)
	if newErr != nil {
		generatedErr := utils.GenerateError("Error encoding JSON", newErr)
		utils.LogError.Print(generatedErr)
		response.Message = fmt.Sprint(generatedErr)
		h.WriteHeader(422)
		json.NewEncoder(h).Encode(response)
	}
}
