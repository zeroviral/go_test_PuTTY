package main

import (
	"github.com/gorilla/mux"
	"go_test_PuTTY/api"
	"go_test_PuTTY/config"
	"go_test_PuTTY/utils"
	_ "go_test_PuTTY/utils"
	"net/http"
)

func main() {
	utils.Init()
	cfg := config.CreateConfiguration()
	router := mux.NewRouter().StrictSlash(true)
	api.RegisterRoutes(router)
	router.Use(utils.LogRequestMetaData)
	utils.LogInfo.Printf("Started %s version %.2f successfully. Running on port %s...", cfg.App.Name, cfg.App.Version, cfg.Server.Port)
	// 06/27/2020 @Travis: Log all requests from -> Server started at whatever port. For now we'll use localhost.
	// Uses router declaration as the "servlet" equivalent
	utils.LogError.Println(http.ListenAndServe(cfg.Server.Port, router))
}
