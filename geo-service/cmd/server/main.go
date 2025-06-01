package main

import (
	"geo/internal/client"
	"geo/internal/config"
	_ "geo/internal/docs"
	"geo/internal/handler"
	"geo/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

func main() {
	log.Println("🟢 geo-service is starting...")
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	cfg := config.NewConfig()
	DadataClient := client.NewDaDataClient(cfg.BaseURL, cfg.Timeout, cfg.Token)
	addressService := service.NewAddressService(DadataClient)
	addressHandler := handler.NewAddressHandler(addressService)

	r.Post("/api/address/search", addressHandler.Search)
	r.Post("/api/address/geocode", addressHandler.Geocode)
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	http.ListenAndServe(":8081", r)
}
