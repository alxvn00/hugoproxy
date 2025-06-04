package auth

import (
	"encoding/json"
	"geo/internal/model"
	"net/http"
)

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
