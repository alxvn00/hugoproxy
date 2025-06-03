package auth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthService_Register_Success(t *testing.T) {
	store := NewMockUserStore()
	issuer := &MockTokenIssuer{}
	authService := NewAuthService(store, issuer)

	err := authService.Register("user@example.com", "password123")
	assert.NoError(t, err)

	_, err = store.Get("user@example.com")
	assert.NoError(t, err)
}

func TestAuthService_Register_AlreadyExists(t *testing.T) {
	store := NewMockUserStore()
	issuer := &MockTokenIssuer{}
	authService := NewAuthService(store, issuer)

	_ = authService.Register("user@example.com", "password123")
	err := authService.Register("user@example.com", "anotherpass")

	assert.EqualError(t, err, "user with email already exists")
}

func TestAuthService_Login_Success(t *testing.T) {
	store := NewMockUserStore()
	issuer := &MockTokenIssuer{token: "test-token"}
	authService := NewAuthService(store, issuer)

	_ = authService.Register("user@example.com", "password123")

	token, err := authService.Login("user@example.com", "password123")
	assert.NoError(t, err)
	assert.Equal(t, "test-token", token)
}

func TestAuthService_Login_WrongPassword(t *testing.T) {
	store := NewMockUserStore()
	issuer := &MockTokenIssuer{}
	authService := NewAuthService(store, issuer)

	_ = authService.Register("user@example.com", "password123")

	_, err := authService.Login("user@example.com", "wrongpassword")
	assert.EqualError(t, err, "invalid password")
}

func TestAuthService_Login_UserNotFound(t *testing.T) {
	store := NewMockUserStore()
	issuer := &MockTokenIssuer{}
	authService := NewAuthService(store, issuer)

	_, err := authService.Login("ghost@example.com", "irrelevant")
	assert.EqualError(t, err, "user not found")
}
