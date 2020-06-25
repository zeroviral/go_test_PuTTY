package endpoints

import (
	"fmt"
	"net/http"
)

func HomeLink(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "hello\n")
	fmt.Fprint(w, "Welcome to the HomePage.")
}
