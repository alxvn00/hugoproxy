package address

import (
	"geo/internal/model"
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
func dummySearch(w http.ResponseWriter, r *http.Request) {}

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
func dummyGeocode(w http.ResponseWriter, r *http.Request) {}

var _ = model.RegisterRequest{}
