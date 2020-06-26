package main

import (
	"github.com/gorilla/mux"
	"go_test_PuTTY/endpoints"
	"log"
	"net/http"
	"os"
)

var LogDir = "resources/logs/info.log"

func main() {

	// Logger file output stuff.
	file, err := os.OpenFile(LogDir, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
	defer file.Close()

	//log.SetOutput(file)

	router := mux.NewRouter().StrictSlash(true)

	// Add our endpoints here.
	router.HandleFunc("/", endpoints.Index)
	router.HandleFunc("/get_transaction", endpoints.GetTransaction)
	router.HandleFunc("/test_endpoint", endpoints.GetTransaction)

	log.Print("Started PuTTY server [ SUCCESSFULLY ]")
	log.Print("Now running PuTTY on 8080...")

	// @Travis: Log all requests from -> Server started at port 8080, localhost.
	// Uses router declaration as the "servlet" equivalent
	log.Fatal(http.ListenAndServe(":8080", router))
}
