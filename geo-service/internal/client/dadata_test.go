package client

import (
	"encoding/json"
	"geo/internal/model"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGeocodeDaData_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/geolocate/address" {
			t.Errorf("unexpected URL path: %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("unexpected method: %s", r.Method)
		}

		var req model.RequestAddressGeocode
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			t.Errorf("failed to decode request: %v", err)
		}

		if req.Lat != 55.7558 || req.Lon != 37.6173 {
			t.Errorf("unexpected coords: %+v", req)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"suggestions": []map[string]interface{}{
				{
					"value":              "Москва",
					"unrestricted_value": "г Москва, Россия",
				},
			},
		})
	}))
	defer server.Close()

	// инициализируем клиента
	client := &DaDataClient{
		BaseURL: server.URL,
		Timeout: 2 * time.Second,
		Token:   "fake-token",
		HTTP:    server.Client(), // используем тот же http.Client
	}

	addresses, err := client.GeocodeDaData(55.7558, 37.6173)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(addresses) != 1 || addresses[0].Value != "Москва" {
		t.Errorf("unexpected response: %+v", addresses)
	}
}
