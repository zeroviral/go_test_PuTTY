package validators

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"go_test_PuTTY/resources"
	"go_test_PuTTY/utils"
	"net/http"
	"strings"
)

// TODO: Decouple JSON parse validation logic with schema validation.
func ValidateRequest(req *http.Request) (resources.EthereumRequest, error) {
	// JSON parsing Validation
	var validRequest resources.EthereumRequest
	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&validRequest)
	if err != nil {
		return validRequest, fmt.Errorf("incorrect request format. please check your inputs")
	}
	// Schema enforcement validation
	if _, err = ValidateSchema(validRequest); err != nil  {
		return validRequest, err
	}
	return validRequest, nil
}

func ValidateSchema(validRequest interface{}) (bool, error) {

	validate := validator.New()
	err := validate.Struct(validRequest)
	if err != nil {
		var errors strings.Builder
		for _, err := range err.(validator.ValidationErrors) {
			utils.LogError.Printf("%v\n", err)
			errors.WriteString(err.Field())
		}
		return false, fmt.Errorf("error on schema validation %s", errors.String())
	}
		return true, nil
}

// Validates a request is of a valid expected action.
func ValidateRequestAction(req *http.Request, allowedAction string) bool {
	if strings.ToLower(req.Method) != strings.ToLower(allowedAction) {
		return false
	}
	return true
}
