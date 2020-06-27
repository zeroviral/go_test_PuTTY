package controllers

import (
	"go_test_PuTTY/utils"
	"log"
	"net/http"
)

type RootHandler func(http.ResponseWriter, *http.Request) error

// This root handler wraps all handlers to provide global error handling logic.
func (fn RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn(w, r)
	if err == nil {
		return
	}
	log.Printf("An error occured: %v", err)
	error4XX, ok := err.(utils.IError4XX)
	if !ok {
		// 500 internal server error
		w.WriteHeader(500)
		return
	}
	body, err := error4XX.ResponseBody()
	if err != nil {
		log.Printf("An error occured internally: %v", err)
		w.WriteHeader(500)
		return
	}
	status, headers := error4XX.ResponseHeaders() // Get http status code and headers.
	for k, v := range headers {
		w.Header().Set(k, v)
	}
	w.WriteHeader(status)
	w.Write(body)
}
