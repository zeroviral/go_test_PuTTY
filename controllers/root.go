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
	clientError, ok := err.(utils.Error4XX)
	if !ok {
		// 500 internal server error
		w.WriteHeader(500)
		return
	}
	body, err := clientError.ResponseBody()
	if err != nil {
		log.Printf("An error occured internally: %v", err)
		w.WriteHeader(500)
		return
	}
	status, headers := clientError.ResponseHeaders() // Get http status code and headers.
	for k, v := range headers {
		w.Header().Set(k, v)
	}
	w.WriteHeader(status)
	w.Write(body)
}
