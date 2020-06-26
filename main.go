package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go_test_PuTTY/endpoints"
	"log"
	"net/http"
	"sync"
)

var once sync.Once

type singleton map[string] string

var (
	ourInstance singleton
)

func connect() singleton {
	once.Do(func() {
		ourInstance = make(singleton)
	})
		return ourInstance
}

func main() {

	//connectionURL := "https://mainnet.infura.io/v3/b1e85ed87262471ea8ce2fd3e6d2f7c8"

	// Connect to ethereum network through infura
	//ourClient := utils.ConnectClient(connectionURL)

	router := mux.NewRouter().StrictSlash(true)

	//_ = ourClient
	// Add our endpoints here.
	router.HandleFunc("/", endpoints.Index)
	router.HandleFunc("/get_transaction", endpoints.GetTransaction)
	router.HandleFunc("/test_endpoint", endpoints.GetTransaction)


	fmt.Println("Congrats, the build ran Successfully!")
	fmt.Println("Now running server on 8080...")

	// @Travis: Log all requests from -> Server started at port 8080, localhost.
	// Uses router declaration as the "servlet" equivalent
	log.Fatal(http.ListenAndServe(":8080", router))
}
