package api

import (
	"github.com/gorilla/mux"
	"go_test_PuTTY/controllers"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", controllers.HealthCheck)
	router.HandleFunc("/transactions", controllers.GetCurrentEthBalance)
	router.HandleFunc("/test_endpoint", controllers.GetCurrentEthBalance)
}
