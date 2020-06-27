package api

import (
	"github.com/gorilla/mux"
	"go_test_PuTTY/controllers"
	"go_test_PuTTY/utils"
	"log"
	"net/http"
)

type rootHandler func(http.ResponseWriter, *http.Request) error

func (fn rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", controllers.HealthCheck)
	router.Handle("/transactions", rootHandler(controllers.GetCurrentEthBalance))
	//router.HandleFunc("/test_endpoint", controllers.GetCurrentEthBalance)
}
