package model

type IncomingGeocodeRequest struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type RequestAddressGeocode struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}
