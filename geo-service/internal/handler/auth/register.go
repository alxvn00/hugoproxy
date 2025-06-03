package auth

import (
	"encoding/json"
	"geo/internal/model"
	"net/http"
)

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

func (h *AuthHandlerImpl) Register(w http.ResponseWriter, r *http.Request) {
	var req model.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	err := h.service.Register(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("registered"))
}
