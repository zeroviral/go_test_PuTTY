package resources

// This will conform our request structure.
type EthereumRequest struct {

	// The field that will be accepted on the left, and the right is the response formatted fields.
	EthereumId string `json:"ethereum_ID"`
}

// This holds our definition of a list that can hold multiple ethereum requests.
type EthereumRequestList []EthereumRequest

type EthereumResponse struct {

	EthereumId string `json:"ethereum_ID"`
	Message    string `json:"message"`
}