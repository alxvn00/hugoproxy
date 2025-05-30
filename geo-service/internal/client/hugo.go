package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"geo/internal/model"
	"net/http"
	"time"
)

type HugoClient struct {
	BaseURL string
	Timeout time.Duration
	HTTP    *http.Client
}

func NewHugoClient(baseURL string, timeout time.Duration) *HugoClient {
	return &HugoClient{
		BaseURL: baseURL,
		Timeout: timeout,
		HTTP:    &http.Client{Timeout: timeout},
	}
}

func (c *HugoClient) Search(query string) ([]*model.Address, error) {
	reqBody := model.RequestAddressSearch{Query: query}
	data, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	url := c.BaseURL + "/address/search"

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status code %d", resp.StatusCode)
	}

	var hugoResp model.ResponseAddress
	err = json.NewDecoder(resp.Body).Decode(&hugoResp)
	if err != nil {
		return nil, err
	}

	return hugoResp.Addresses, nil
}

func (c *HugoClient) Geocode(lat, lon float64) ([]*model.Address, error) {
	reqBody := model.RequestAddressGeocode{Lat: lat, Lon: lon}
	data, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	url := c.BaseURL + "/address/geocode"

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status code %d", resp.StatusCode)
	}

	var addresses []*model.Address
	err = json.NewDecoder(resp.Body).Decode(&addresses)
	if err != nil {
		return nil, err
	}

	return addresses, nil
}
