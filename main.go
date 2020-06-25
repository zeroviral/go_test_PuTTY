package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "hello\n")
	fmt.Fprint(w, "Welcome to the HomePage.")
}

// ENDPOINT: ("/get_transaction")
// ACCEPTS: GET
func getTransAction(w http.ResponseWriter, r *http.Request) {
	// TODO: Fill in the code for sourcing transaction from Ethereum

	// blah blah blah
	w.WriteHeader(200)
}

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatalf("Unable to connect")
	}
	_ = client
	fmt.Println("Successfully connected to Infura/Ethereum")

	router := mux.NewRouter().StrictSlash(true)

	// @Travis: Declare what our "/" endpoint will do.
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/get_transaction", getTransAction)
	fmt.Println("Congrats, the build ran Successfully!")
	fmt.Println("Now running server on 8080...")

	// @Travis: Log all requests from -> Server started at port 8080, localhost.
	// Uses router declaration as the "servlet" equivalent
	log.Fatal(http.ListenAndServe(":8080", router))
}
