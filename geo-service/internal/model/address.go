package model

type RequestAddressSearch struct {
	Query string `json:"query"`
}

type Address struct {
	Value             string       `json:"value"`
	UnrestrictedValue string       `json:"unrestricted_value"`
	Data              *AddressData `json:"data"`
}

type AddressData struct {
	City    string `json:"city,omitempty"`
	Street  string `json:"street,omitempty"`
	House   string `json:"house,omitempty"`
	GeoLat  string `json:"geo_lat,omitempty"`
	GeoLong string `json:"geo_long,omitempty"`
}

type ResponseAddress struct {
	Addresses []*Address `json:"addresses"`
}
