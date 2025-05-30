package main

import (
	"geo/internal/client"
	"geo/internal/config"
	"geo/internal/handler"
	"geo/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	cfg := config.NewConfig()
	hugoClient := client.NewHugoClient(cfg.BaseURL, cfg.Timeout)
	addressService := service.NewAddressService(hugoClient)
	addressHandler := handler.NewAddressHandler(addressService)

	r.HandleFunc("/api/address/search", addressHandler.Search)
	r.HandleFunc("/address/geocode", addressHandler.Geocode)

	http.ListenAndServe(":8081", r)
}
