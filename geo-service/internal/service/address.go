package service

import (
	"geo/internal/client"
	"geo/internal/model"
)

type AddressService struct {
	Client *client.HugoClient
}

func NewAddressService(client *client.HugoClient) *AddressService {
	return &AddressService{Client: client}
}

func (s *AddressService) Search(query string) ([]*model.Address, error) {
	return s.Client.Search(query)
}

func (s *AddressService) Geocode(lat, lng float64) ([]*model.Address, error) {
	return s.Client.Geocode(lat, lng)
}
