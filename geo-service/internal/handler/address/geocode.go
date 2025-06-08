package address

import (
	"encoding/json"
	"fmt"
	"github.com/alxvn00/hugoproxy/geo-service/internal/model"
	"log"
	"net/http"
)

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
