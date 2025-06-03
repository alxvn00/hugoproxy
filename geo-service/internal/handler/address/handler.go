package address

import (
	"geo/internal/service/address"
	"net/http"
)

type AddressHandler interface {
	Search(w http.ResponseWriter, r *http.Request)
	Geocode(w http.ResponseWriter, r *http.Request)
}

type AddressHandlerImpl struct {
	Service address.AddressService
}

func NewAddressHandler(service address.AddressService) *AddressHandlerImpl {
	return &AddressHandlerImpl{Service: service}
}
