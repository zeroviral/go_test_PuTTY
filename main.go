package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the HomePage.")
}

func main() {
	fmt.Println("Congrats, the build ran Successfully!")
	fmt.Println("Now running server on 8080...")

	router := mux.NewRouter().StrictSlash(true)

	// @Travis: Declare what our "/" endpoint will do.
	router.HandleFunc("/", homeLink)

	// @Travis: Log all requests from -> Server started at port 8080, localhost.
	// Uses router declaration as the "servlet" equivalent
	log.Fatal(http.ListenAndServe(":8080", router))
}
