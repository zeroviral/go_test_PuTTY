package resources

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"go_test_PuTTY/utils"
	"net/http"
	"strings"
)

// This will conform our request structure.
type TransactionsSchemaGET struct {

	// The field that will be accepted on the left, and the right is the response formatted fields.
	EID string `json:"eID" validate:"required,gt=38,lt=43"`
}

// This holds our definition of a list that can hold multiple ethereum requests.
type EthereumRequestList []TransactionsSchemaGET

type EthereumResponse struct {
	EthereumId string `json:"eID"`
	Message    string `json:"message"`
}


func (e *TransactionsSchemaGET) ValidateRequest (req *http.Request) error {
	if err := e.ValidateJSON(req); err != nil {
		return err
	}
	if err := e.ValidateSchema(); err != nil {
		return err
	}
	return nil
}

func (e *TransactionsSchemaGET) ValidateJSON (req *http.Request) error {
	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(e)
	if err != nil {
		return fmt.Errorf("Incorrect JSON format")
	}
	return nil
}

func (e *TransactionsSchemaGET) ValidateSchema () error {
	validate := validator.New()
	err := validate.Struct(e)
	if err != nil {
		var errors strings.Builder
		for _, err := range err.(validator.ValidationErrors) {
			utils.LogError.Printf("%v\n", err)
			errors.WriteString(err.Field())
		}
		return fmt.Errorf("error on schema validation %s", errors.String())
	}
	return nil
}
