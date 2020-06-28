package validators

import (
	"fmt"
	"go_test_PuTTY/resources"
	"net/http"
	"strings"
)

// TODO: Decouple JSON parse validation logic with schema validation.
func ValidateRequest(req *http.Request) (resources.EthereumRequest, error) {
	// JSON parsing Validation
	var ethereumRequest resources.EthereumRequest
	if err := ethereumRequest.ValidateJSON(req); err != nil {
		return ethereumRequest, fmt.Errorf("incorrect request format. please check your inputs")
	}
	// Schema enforcement validation
	if err := ethereumRequest.ValidateSchema(); err != nil  {
		return ethereumRequest, err
	}
	return ethereumRequest, nil
}

// Validates a request is of a valid expected action.
func ValidateRequestAction(req *http.Request, allowedAction string) bool {
	if strings.ToLower(req.Method) != strings.ToLower(allowedAction) {
		return false
	}
	return true
}
