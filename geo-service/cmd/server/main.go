package main

import (
	"geo/internal/app"
	"geo/internal/config"
	_ "geo/internal/docs"
	_ "geo/internal/model"
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
