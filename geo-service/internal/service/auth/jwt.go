package auth

import (
	"github.com/go-chi/jwtauth/v5"
	"time"
)

type TokenIssuer interface {
	Issue(userEmail string) (string, error)
}

type JWTManager struct {
	Auth *jwtauth.JWTAuth
	TTL  time.Duration
}

func NewJWTManager(secret string, ttl time.Duration) *JWTManager {
	return &JWTManager{
		Auth: jwtauth.New("HS256", []byte(secret), nil),
		TTL:  ttl,
	}
}

func (j *JWTManager) Issue(userEmail string) (string, error) {
	_, tokenStr, err := j.Auth.Encode(map[string]interface{}{
		"user_email": userEmail,
		"exp":        time.Now().Add(j.TTL).Unix(),
		"iat":        time.Now().Unix(),
	})
	return tokenStr, err
}
