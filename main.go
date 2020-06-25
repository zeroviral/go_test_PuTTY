package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go_test_PuTTY/endpoints"
	"go_test_PuTTY/utils"
	"log"
	"net/http"
)

func main() {

	connectionURL := "https://mainnet.infura.io/v3/b1e85ed87262471ea8ce2fd3e6d2f7c8"

	// Connect to ethereum network through infura
	utils.ConnectClient(connectionURL)

	router := mux.NewRouter().StrictSlash(true)

	// @Travis: Declare what our "/" endpoint will do.
	router.HandleFunc("/", endpoints.HomeLink)
	router.HandleFunc("/get_transaction", endpoints.GetTransAction)
	fmt.Println("Congrats, the build ran Successfully!")
	fmt.Println("Now running server on 8080...")

	// @Travis: Log all requests from -> Server started at port 8080, localhost.
	// Uses router declaration as the "servlet" equivalent
	log.Fatal(http.ListenAndServe(":8080", router))
}
