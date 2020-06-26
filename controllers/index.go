package controllers

import (
	"fmt"
	"net/http"
)

// Serve a simple yet effective index.
func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Welcome to the HomePage.")
}
