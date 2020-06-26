package api

import (
	"github.com/gorilla/mux"
	"go_test_PuTTY/controllers"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", controllers.Index)
	router.HandleFunc("/transactions", controllers.GetTransaction)
	router.HandleFunc("/test_endpoint", controllers.GetTransaction)
}
