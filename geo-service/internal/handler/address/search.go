package address

import (
	"encoding/json"
	"geo/internal/model"
	"net/http"
	"strings"
)

// Search
// @Summary Поиск адреса по тексту
// @Description Автодополнение адреса через DaData по текстовому запросу
// @Tags address
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param input body model.RequestAddressSearch true "Поисковый запрос"
// @Success 200 {object} model.ResponseAddress
// @Failure 400 {string} string "Bad Request"
// @Failure 403 {string} string "Forbidden"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/address/search [post]

func (h *AddressHandlerImpl) Search(w http.ResponseWriter, r *http.Request) {
	var req model.RequestAddressSearch

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := strings.TrimSpace(req.Query)
	if query == "" {
		http.Error(w, "invalid query", http.StatusBadRequest)
		return
	}

	addresses, err := h.Service.Search(query)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.ResponseAddress{
		Addresses: addresses,
	})
}
