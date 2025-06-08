package auth

import (
	"github.com/alxvn00/hugoproxy/geo-service/internal/model"
	"github.com/alxvn00/hugoproxy/geo-service/internal/service/auth"
	"net/http"
)

type AuthHandler interface {
	Register(http.ResponseWriter, *http.Request)
	Login(http.ResponseWriter, *http.Request)
}

type AuthHandlerImpl struct {
	service auth.AuthService
}

func NewAuthHandler(service auth.AuthService) AuthHandler {
	return &AuthHandlerImpl{service: service}
}

// Login godoc
// @Summary      Вход пользователя
// @Description  Аутентифицирует пользователя и возвращает JWT
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body model.LoginRequest true "Email и пароль"
// @Success      200 {object} model.TokenResponse
// @Failure      200 {string} string "Invalid credentials"
// @Failure      500 {string} string "Internal error"
// @Router       /api/login [post]
func dummyLogin(w http.ResponseWriter, r *http.Request) {}

// Register godoc
// @Summary      Регистрация пользователя
// @Description  Регистрирует нового пользователя
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body model.RegisterRequest true "Email и пароль"
// @Success      200 {string} string "OK"
// @Failure      400 {string} string "Bad Request"
// @Failure      500 {string} string "Internal Error"
// @Router       /api/register [post]
func dummyRegister(w http.ResponseWriter, r *http.Request) {}

var _ = model.RegisterRequest{}
