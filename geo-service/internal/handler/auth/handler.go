package auth

import (
	"geo/internal/service/auth"
	"net/http"
)

type AuthHandler interface {
	Register(http.ResponseWriter, *http.Request)
	Login(http.ResponseWriter, *http.Request)
}

type AuthHandlerImpl struct {
	service *auth.AuthService
}

func NewAuthHandler(service *auth.AuthService) AuthHandler {
	return &AuthHandlerImpl{service: service}
}
