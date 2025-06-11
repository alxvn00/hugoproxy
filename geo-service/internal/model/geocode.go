package model

type IncomingGeocodeRequest struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type RequestAddressGeocode struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}
