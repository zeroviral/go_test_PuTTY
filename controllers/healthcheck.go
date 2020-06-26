package controllers

import (
	"encoding/json"
	"fmt"
	"go_test_PuTTY/config"
	"net/http"
)

type HealthCheckResponse struct {
	Name    string  `json:"name"`
	Version float64 `json:"version"`
	Commit  string  `json:"commit"`
	Host    string  `json:"host"`
}

// HealthCheck Endpoint
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	cfg := config.CreateConfiguration()
	resp := HealthCheckResponse{
		Name:    cfg.App.Name,
		Version: cfg.App.Version,
		Commit:  "Commit hash",
		Host:    cfg.Server.Host,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Printf("%v", resp)
	//respJSON, err := json.Marshal(resp)
	//if err != nil {
	//	// handle error
	//}
	json.NewEncoder(w).Encode(resp)
}
