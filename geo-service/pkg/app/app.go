package app

import (
	"geo/internal/client"
	"geo/internal/handler/address"
	authHandler "geo/internal/handler/auth"
	address2 "geo/internal/service/address"
	"geo/internal/service/auth"
	"geo/pkg/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"time"
)

func Init(cfg *config.Config) *chi.Mux {
	userStore := auth.NewMemoryUserStore()
	jwtManager := auth.NewJWTManager(cfg.JWTSecret, time.Hour*2)
	authService := auth.NewAuthService(userStore, jwtManager)
	authHandler := authHandler.NewAuthHandler(authService)

	dadataClient := client.NewDaDataClient(cfg.BaseURL, cfg.Timeout, cfg.Token)
	addressService := address2.NewAddressService(dadataClient)
	addressHandler := address.NewAddressHandler(addressService)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/api/register", authHandler.Register)
	r.Post("/api/login", authHandler.Login)

	r.Group(func(protected chi.Router) {
		protected.Use(jwtauth.Verifier(jwtManager.Auth))
		protected.Use(jwtauth.Authenticator(jwtManager.Auth))

		protected.Post("/api/address/search", addressHandler.Search)
		protected.Post("/api/address/geocode", addressHandler.Geocode)
	})

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	return r
}
