package controllers

import (
	"encoding/json"
	"fmt"
	"go_test_PuTTY/config"
	"net/http"
)

type HealthCheckResponse struct {
	name    string
	version float64
	commit  string
	host    string
}

// HealthCheck Endpoint
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	cfg := config.CreateConfiguration()
	resp := HealthCheckResponse{
		name:    cfg.App.Name,
		version: cfg.App.Version,
		commit:  "Commit hash",
		host:    cfg.Server.Host,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Printf("%v", resp)
	respJSON, err := json.Marshal(resp)
	if err != nil {
		// handle error
	}
	w.Write(respJSON)
}
