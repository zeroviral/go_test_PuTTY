package validators

import (
	"encoding/json"
	"go_test_PuTTY/resources"
	"go_test_PuTTY/utils"
	"io/ioutil"
	"net/http"
	"strings"
)

func ValidateRequest(req *http.Request) (resources.EthereumRequest, error) {
	var validRequest resources.EthereumRequest
	body, _ := ioutil.ReadAll(req.Body)
	if err := json.Unmarshal(body, &validRequest); err != nil {
		utils.LogError.Println("Invalid Request Format. Unable to unmarshal JSON.")
		return validRequest, err
	}
	return validRequest, nil
}

// Validates a request is of a valid expected action.
func ValidateRequestAction(w http.ResponseWriter, req *http.Request, allowedAction string) bool {
	if strings.ToLower(req.Method) != strings.ToLower(allowedAction) {
		utils.LogError.Printf("ACTION NOT ALLOWED. RECEIVED: %s", req.Method)
		http.Error(w, `Method was `+req.Method+`, expected ` + allowedAction, http.StatusNotFound)
		return false
	}
	return true
}