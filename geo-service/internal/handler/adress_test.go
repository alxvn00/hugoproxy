package handler

import (
	"encoding/json"
	"geo/internal/model"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type mockAddressService struct {
	GeocodeFn func(lat, lng float64) ([]*model.Address, error)
	SearchFn  func(query string) ([]*model.Address, error)
}

func (m *mockAddressService) Geocode(lat, lng float64) ([]*model.Address, error) {
	return m.GeocodeFn(lat, lng)
}

func (m *mockAddressService) Search(query string) ([]*model.Address, error) {
	return m.SearchFn(query)
}

func TestGeocodeHandler_Success(t *testing.T) {
	mockService := &mockAddressService{
		GeocodeFn: func(lat, lng float64) ([]*model.Address, error) {
			return []*model.Address{
				{Value: "Москва", UnrestrictedValue: "г Москва, Россия"},
			}, nil
		},
		SearchFn: func(query string) ([]*model.Address, error) {
			return nil, nil
		},
	}

	handler := NewAddressHandler(mockService)

	reqBody := `{"lat":55.7558,"lng":37.6173}`
	req := httptest.NewRequest(http.MethodPost, "/api/address/geocode", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	handler.Geocode(rec, req)

	resp := rec.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}

	var body model.ResponseAddress
	err := json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(body.Addresses) != 1 || body.Addresses[0].Value != "Москва" {
		t.Errorf("unexpected response: %+v", body)
	}
}

func TestSearchHandler_Success(t *testing.T) {
	mockService := &mockAddressService{
		SearchFn: func(query string) ([]*model.Address, error) {
			return []*model.Address{
				{Value: "Санкт-Петербург", UnrestrictedValue: "г Санкт-Петербург, Россия"},
			}, nil
		},
		GeocodeFn: func(lat, lng float64) ([]*model.Address, error) {
			return nil, nil
		},
	}

	handler := NewAddressHandler(mockService)

	reqBody := `{"query":"Санкт-Петербург"}`
	req := httptest.NewRequest(http.MethodPost, "/api/address/search", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	handler.Search(rec, req)

	resp := rec.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}

	var body model.ResponseAddress
	err := json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(body.Addresses) != 1 || body.Addresses[0].Value != "Санкт-Петербург" {
		t.Errorf("unexpected response: %+v", body)
	}
}
