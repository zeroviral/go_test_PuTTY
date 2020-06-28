package api

import (
	"github.com/gorilla/mux"
	ctrl "go_test_PuTTY/controllers"
)

func RegisterRoutes(router *mux.Router) {
	router.Handle("/health", ctrl.RootHandler(ctrl.HealthCheck))
	router.Handle("/transactions", ctrl.RootHandler(ctrl.TransactionsHandler))
}
