package endpoints

import (
	"fmt"
	"net/http"
)

// ENDPOINT: ("/get_transaction")
// ACCEPTS: GET
func GetTransAction(w http.ResponseWriter, r *http.Request) {
	// TODO: Fill in the code for sourcing transaction from Ethereum

	switch r.Method {

	case "GET":

		payload := r.URL.Query()
		fmt.Println(payload)

		fmt.Println(payload.Get("somefield"))
		//ourQuery := r.URL.Query()

	}
}
