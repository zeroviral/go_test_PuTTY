package validators

import (
	"fmt"
	"go_test_PuTTY/resources"
	"net/http"
)

// Validate GET requests on /transactions
func TransactionsGET(req *http.Request) (resources.TransactionsSchemaGET, error) {
	var request resources.TransactionsSchemaGET
	if err := request.ValidateRequest(req); err != nil {
		return request, fmt.Errorf("incorrect request format. please check your inputs")
	}
	return request, nil
}
