package controllers

import (
	"go_test_PuTTY/utils"
	"net/http"
)

type RootHandler func(http.ResponseWriter, *http.Request) error

// This root handler wraps all handlers to provide global error handling logic.
func (fn RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn(w, r)
	if err == nil {
		return
	}
	error4XX, ok := err.(utils.IError4XX)
	if !ok {
		utils.LogError.Printf("An error occured internally: %v", err)
		// ONLY set the status code if the header is not set.
		if w.Header() == nil {
			// 500 internal server error
			w.WriteHeader(500)
		}
		return
	}
	body, err := error4XX.ResponseBody()
	if err != nil {
		utils.LogError.Printf("An error occured internally: %v", err)
		if w.Header() == nil {
			w.WriteHeader(500)
		}
		return
	}
	status, headers := error4XX.ResponseHeaders() // Get http status code and headers.
	for k, v := range headers {
		w.Header().Set(k, v)
	}
	utils.LogError.Printf("Code %v %s %s %v txID:[%s]", status, r.Method, r.RequestURI, string(body), r.Context().Value("transactionID"))
	w.WriteHeader(status)
	w.Write(body)
}
