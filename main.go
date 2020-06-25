package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
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
	fmt.Println("Congrats, the build ran Successfully!")
	fmt.Println("Now running server on 8080...")

	router := mux.NewRouter().StrictSlash(true)

	// @Travis: Declare what our "/" endpoint will do.
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/get_transaction", getTransAction)

	// @Travis: Log all requests from -> Server started at port 8080, localhost.
	// Uses router declaration as the "servlet" equivalent
	log.Fatal(http.ListenAndServe(":8080", router))
}
