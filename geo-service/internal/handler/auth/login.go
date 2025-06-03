package auth

import (
	"encoding/json"
	"geo/internal/model"
	"net/http"
)

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

func (h *AuthHandlerImpl) Login(w http.ResponseWriter, r *http.Request) {
	var req model.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	token, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(model.TokenResponse{Token: token})
}
