package handler

import (
	"encoding/json"
	"fmt"
	"geo/internal/model"
	"geo/internal/service"
	"log"
	"net/http"
	"strings"
)

type AddressHandler struct {
	Service service.AddressService
}

func NewAddressHandler(service service.AddressService) *AddressHandler {
	return &AddressHandler{Service: service}
}

// Search
// @Summary Поиск адреса по тексту
// @Description Автодополнение адреса через DaData по текстовому запросу
// @Tags address
// @Accept json
// @Produce json
// @Param input body model.RequestAddressSearch true "Поисковый запрос"
// @Success 200 {object} model.ResponseAddress
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/address/search [post]
func (h *AddressHandler) Search(w http.ResponseWriter, r *http.Request) {
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

// Geocode
// @Summary Геокодирование по координатам
// @Description Получение адресов через DaData по lat/lng
// @Tags address
// @Accept json
// @Produce json
// @Param input body model.IncomingGeocodeRequest true "Широта и долгота"
// @Success 200 {object} model.ResponseAddress
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/address/geocode [post]
func (h *AddressHandler) Geocode(w http.ResponseWriter, r *http.Request) {
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
