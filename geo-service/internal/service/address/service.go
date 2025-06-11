package address

import (
	"github.com/alxvn00/hugoproxy/geo-service/internal/client"
	"github.com/alxvn00/hugoproxy/geo-service/internal/model"
	"log"
)

type AddressService interface {
	Geocode(lat, lng string) ([]*model.Address, error)
	Search(query string) ([]*model.Address, error)
}

type addressService struct {
	client *client.DaDataClient
}

var _ AddressService = (*addressService)(nil)

func NewAddressService(client *client.DaDataClient) AddressService {
	return &addressService{client: client}
}

func (s *addressService) Search(query string) ([]*model.Address, error) {
	log.Printf("📥 [Search] query: %s", query)
	return s.client.SearchDaData(query)
}

func (s *addressService) Geocode(lat, lng string) ([]*model.Address, error) {
	log.Printf("📥 [Geocode] lat=%.6f, lon=%.6f", lat, lng)
	return s.client.GeocodeDaData(lat, lng)
}
