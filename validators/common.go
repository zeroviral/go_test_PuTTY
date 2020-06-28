package validators

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"go_test_PuTTY/resources"
	"go_test_PuTTY/utils"
	"io/ioutil"
	"net/http"
	"strings"
)

// TODO: move all validation logic to single aggregator function
func ValidateRequest(req *http.Request) (resources.EthereumRequest, error) {
	var validRequest resources.EthereumRequest
	body, _ := ioutil.ReadAll(req.Body)
	if err := json.Unmarshal(body, &validRequest); err != nil {
		utils.LogError.Println("Invalid Request Format. Unable to unmarshal JSON.")
		return validRequest, err
	}
	validate := validator.New()
	err := validate.Struct(validRequest)
	if err != nil {
		var errors strings.Builder
		for _, err := range err.(validator.ValidationErrors) {
			utils.LogError.Printf("%v\n", err)
			errors.WriteString(err.Field() + ":" + err.ActualTag() + ", ")
		}
		return validRequest, fmt.Errorf("Failed to validate on the following fields: %s", errors.String())
	}
	return validRequest, nil
}

// Validates a request is of a valid expected action.
func ValidateRequestAction(req *http.Request, allowedAction string) bool {
	if strings.ToLower(req.Method) != strings.ToLower(allowedAction) {
		return false
	}
	return true
}

func ValidateSchema(req *http.Request, schema interface{}) {
	return
}
