package model

type IncomingGeocodeRequest struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type RequestAddressGeocode struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}
