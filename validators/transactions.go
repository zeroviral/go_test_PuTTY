package validators

import "net/http"

// GET request validator for transactions
func ValidateTxGET(req *http.Request) {
	ValidateRequestAction(req, "GET")
	ValidateSchema(req, 1)
}

// POST request validator for transactions
func ValidateTxPOST(req *http.Request) {
	ValidateRequestAction(req, "POST")
	ValidateSchema(req, 1)
}
