package validators

import (
	"encoding/json"
	"go_test_PuTTY/resources"
	"go_test_PuTTY/utils"
	"io/ioutil"
	"net/http"
)

func ValidateRequest(req *http.Request) (resources.EthereumRequest, error) {
	var validRequest resources.EthereumRequest
	body, _ := ioutil.ReadAll(req.Body)
	if err := json.Unmarshal(body, &validRequest); err != nil {
		utils.LogWarning.Println("Invalid Request Format. Unable to unmarshal JSON.")
		return validRequest, err
	}
	return validRequest, nil
}
