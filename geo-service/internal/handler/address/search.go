package address

import (
	"encoding/json"
	"fmt"
	"github.com/alxvn00/hugoproxy/geo-service/internal/model"
	"net/http"
	"strings"
)

func (h *AddressHandlerImpl) Search(w http.ResponseWriter, r *http.Request) {
	var req model.RequestAddressSearch

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.Responder.Error(w, http.StatusBadRequest, err)
		return
	}

	query := strings.TrimSpace(req.Query)
	if query == "" {
		h.Responder.Error(w, http.StatusBadRequest, errInvalidQuery())
		return
	}

	addresses, err := h.Service.Search(query)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	h.Responder.Success(w, http.StatusOK, model.ResponseAddress{Addresses: addresses})
}

func errInvalidQuery() error {
	return fmt.Errorf("invalid query")
}
