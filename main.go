package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	"go_test_PuTTY/endpoints"
	"log"
	"net/http"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "hello\n")
	fmt.Fprint(w, "Welcome to the HomePage.")
}

func connectClient(url string) {
	finalCLient, err := ethclient.Dial(url)
	if err != nil {
		log.Fatalf("Can't connect")
	}

	_ = finalCLient
	fmt.Println("Successfully connected to Infura/Ethereum")
}

func main() {

	connectionURL := "https://mainnet.infura.io/v3/b1e85ed87262471ea8ce2fd3e6d2f7c8"

	// Connect to ethereum network through infura
	connectClient(connectionURL)

	//account := common.HexToAddress("0x742d35cc6634c0532925a3b844bc454e4438f44e")

	//balance, err := client.BalanceAt(context.Background(), account, nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("The balance of this ethereum address is: %s ether \n", balance)
	router := mux.NewRouter().StrictSlash(true)

	// @Travis: Declare what our "/" endpoint will do.
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/get_transaction", endpoints.GetTransAction)
	fmt.Println("Congrats, the build ran Successfully!")
	fmt.Println("Now running server on 8080...")

	// @Travis: Log all requests from -> Server started at port 8080, localhost.
	// Uses router declaration as the "servlet" equivalent
	log.Fatal(http.ListenAndServe(":8080", router))
}
