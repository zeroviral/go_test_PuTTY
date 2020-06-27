package controllers

import (
	"encoding/json"
	"go_test_PuTTY/config"
	"go_test_PuTTY/utils"
	"net/http"
)

type HealthCheckResponse struct {
	Name    string  `json:"name"`
	Version float64 `json:"version"`
	Commit  string  `json:"commit"`
	Host    string  `json:"host"`
}

// HealthCheck Endpoint
func HealthCheck(w http.ResponseWriter, r *http.Request) error {
	cfg := config.CreateConfiguration()
	resp := HealthCheckResponse{
		Name:    cfg.App.Name,
		Version: cfg.App.Version,
		Commit:  "Commit hash",
		Host:    cfg.Server.Host,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	utils.LogInfo.Printf("%v", resp)
	json.NewEncoder(w).Encode(resp)
	return nil
}
