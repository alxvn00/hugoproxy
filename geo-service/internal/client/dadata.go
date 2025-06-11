package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/alxvn00/hugoproxy/geo-service/internal/model"
	"log"
	"net/http"
	"time"
)

type DaDataClient struct {
	BaseURL string
	Timeout time.Duration
	Token   string
	HTTP    *http.Client
}

func NewDaDataClient(baseURL string, timeout time.Duration, token string) *DaDataClient {
	return &DaDataClient{
		BaseURL: baseURL,
		Timeout: timeout,
		Token:   token,
		HTTP:    &http.Client{Timeout: timeout},
	}
}

func (c *DaDataClient) SearchDaData(query string) ([]*model.Address, error) {
	reqBody := model.RequestAddressSearch{
		Query: query,
	}
	data, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	url := c.BaseURL + "/suggest/address"

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Token "+c.Token)

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}

	log.Printf("🌍 Sending to DaData: %s", url)
	log.Printf("📡 Payload: %s", string(data))
	log.Printf("📦 StatusCode: %d", resp.StatusCode)

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("dadata status: %d", resp.StatusCode)
	}

	var raw struct {
		Suggestions []*model.Address `json:"suggestions"`
	}

	err = json.NewDecoder(resp.Body).Decode(&raw)
	if err != nil {
		log.Printf("❌ Failed to decode response: %v", err)
		return nil, err
	}

	return raw.Suggestions, nil
}

func (c *DaDataClient) GeocodeDaData(lat, lng string) ([]*model.Address, error) {
	log.Printf("📥 [Geocode] lat=%.6f, lng=%.6f", lat, lng)

	reqBody := model.RequestAddressGeocode{Lat: lat, Lon: lng}
	data, err := json.Marshal(reqBody)
	if err != nil {
		log.Printf("❌ Failed to marshal request: %v", err)
		return nil, err
	}

	url := c.BaseURL + "/geolocate/address"

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		log.Printf("❌ Failed to create request: %v", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Token "+c.Token)

	log.Printf("🌍 Sending to DaData: %s", url)
	log.Printf("📡 Payload: %s", string(data))

	resp, err := c.HTTP.Do(req)
	if err != nil {
		log.Printf("❌ HTTP request failed: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	log.Printf("📦 StatusCode: %d", resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status code %d", resp.StatusCode)
	}

	var raw struct {
		Suggestions []*model.Address `json:"suggestions"`
	}

	err = json.NewDecoder(resp.Body).Decode(&raw)
	if err != nil {
		log.Printf("❌ Failed to decode response: %v", err)
		return nil, err
	}

	log.Printf("📤 Received %d suggestions", len(raw.Suggestions))
	return raw.Suggestions, nil
}
