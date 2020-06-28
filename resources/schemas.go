package resources

// This will conform our request structure.
type EthereumRequest struct {

	// The field that will be accepted on the left, and the right is the response formatted fields.
	EID string `json:"eID" validate:"required"`
}

// This holds our definition of a list that can hold multiple ethereum requests.
type EthereumRequestList []EthereumRequest

type EthereumResponse struct {
	EthereumId string `json:"eID"`
	Message    string `json:"message"`
}
