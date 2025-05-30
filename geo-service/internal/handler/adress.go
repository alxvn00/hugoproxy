package handler

import (
	"encoding/json"
	"geo/internal/model"
	"geo/internal/service"
	"net/http"
	"strings"
)

type AddressHandler struct {
	Service *service.AddressService
}

func NewAddressHandler(service *service.AddressService) *AddressHandler {
	return &AddressHandler{Service: service}
}

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

func (h *AddressHandler) Geocode(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	addresses, er := h.Service.Geocode(req.Lat, req.Lng)
	if er != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(addresses)
}
