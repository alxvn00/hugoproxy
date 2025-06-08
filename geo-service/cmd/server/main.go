package main

import (
	_ "github.com/alxvn00/hugoproxy/geo-service/internal/docs"
	_ "github.com/alxvn00/hugoproxy/geo-service/internal/model"
	"github.com/alxvn00/hugoproxy/geo-service/pkg/app"
	"github.com/alxvn00/hugoproxy/geo-service/pkg/config"
	"log"
	"net/http"
)

// @title Geo API
// @version 1.0
// @description Геосервис с JWT авторизацией и обработкой адресов
// @host localhost:8080
// @BasePath /
// @schemes http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	cfg := config.NewConfig()
	r := app.Init(cfg)

	log.Println("🟢 geo-service is starting...")
	if err := http.ListenAndServe(cfg.Port, r); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
