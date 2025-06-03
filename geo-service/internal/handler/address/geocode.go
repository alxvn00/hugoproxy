package address

import (
	"encoding/json"
	"fmt"
	"geo/internal/model"
	"log"
	"net/http"
)

// Geocode
// @Summary Геокодирование по координатам
// @Description Получение адресов через DaData по lat/lng
// @Tags address
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param input body model.IncomingGeocodeRequest true "Широта и долгота"
// @Success 200 {object} model.ResponseAddress
// @Failure 400 {string} string "Bad Request"
// @Failure 403 {string} string "Forbidden"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/address/geocode [post]

func (h *AddressHandlerImpl) Geocode(w http.ResponseWriter, r *http.Request) {
	log.Println("📨 Handler /api/address/geocode received request")
	var req model.IncomingGeocodeRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("📥 Parsed lat=%f, lng=%f\n", req.Lat, req.Lng)

	addresses, er := h.Service.Geocode(req.Lat, req.Lng)
	if er != nil {
		fmt.Printf("❌ service.Geocode() failed: %v\n", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(model.ResponseAddress{Addresses: addresses})
}
