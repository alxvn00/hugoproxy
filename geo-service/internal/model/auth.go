package model

type LoginRequest struct {
	Email    string `json:"email" example:"test@test.com"`
	Password string `json:"password" example:"mysecretpassword"`
}

type RegisterRequest struct {
	Email    string `json:"email" example:"test@test.com"`
	Password string `json:"password" example:"mysecretpassword"`
}

type TokenResponse struct {
	Token string `json:"token" example:"eyJusdKioJI.."`
}
