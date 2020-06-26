package main

import (
	"github.com/gorilla/mux"
	"go_test_PuTTY/endpoints"
	"go_test_PuTTY/utils"
	_ "go_test_PuTTY/utils"
	"net/http"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)

	// Add our endpoints here.
	router.HandleFunc("/", endpoints.Index)
	router.HandleFunc("/get_transaction", endpoints.GetTransaction)
	router.HandleFunc("/test_endpoint", endpoints.GetTransaction)

	utils.LogInfo.Println("Started PuTTY server [ SUCCESSFULLY ]")
	utils.LogInfo.Println("Now running PuTTY on port: 8080...")

	// @Travis: Log all requests from -> Server started at port 8080, localhost.
	// Uses router declaration as the "servlet" equivalent
	utils.LogError.Println(http.ListenAndServe(":8080", router))
}
