package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/alxvn00/hugoproxy/geo-service/internal/model"
	authservice "github.com/alxvn00/hugoproxy/geo-service/internal/service/auth"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockUserStore struct {
	users map[string][]byte
}

func (m *mockUserStore) Get(email string) ([]byte, error) {
	hash, ok := m.users[email]
	if !ok {
		return nil, errors.New("user not found")
	}
	return hash, nil
}

func (m *mockUserStore) Save(email string, hash []byte) error {
	if _, exists := m.users[email]; exists {
		return errors.New("user not found")
	}
	m.users[email] = hash
	return nil
}

type mockTokenIssuer struct {
	token string
	err   error
}

func (m *mockTokenIssuer) Issue(email string) (string, error) {
	return m.token, m.err
}

func TestAuthHandler_Register_Success(t *testing.T) {
	store := &mockUserStore{users: make(map[string][]byte)}
	jwt := &mockTokenIssuer{}
	service := authservice.NewAuthService(store, jwt)
	handler := NewAuthHandler(service)

	body := model.RegisterRequest{
		Email:    "test@example.com",
		Password: "secure123",
	}
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/api/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.Register(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Contains(t, w.Body.String(), "registered")
}

func TestAuthHandler_Register_InvalidJSON(t *testing.T) {
	service := authservice.NewAuthService(&mockUserStore{users: make(map[string][]byte)}, &mockTokenIssuer{})
	handler := NewAuthHandler(service)

	req := httptest.NewRequest(http.MethodPost, "/api/register", bytes.NewBuffer([]byte("bad-json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.Register(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	assert.Contains(t, w.Body.String(), "invalid json")
}

func TestAuthHandler_Login_Success(t *testing.T) {
	store := &mockUserStore{users: make(map[string][]byte)}
	hashed, _ := bcrypt.GenerateFromPassword([]byte("secure123"), bcrypt.DefaultCost)
	store.users["user@example.com"] = hashed

	jwt := &mockTokenIssuer{token: "valid-token"}
	service := authservice.NewAuthService(store, jwt)
	handler := NewAuthHandler(service)

	body := model.LoginRequest{
		Email:    "user@example.com",
		Password: "secure123",
	}
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/api/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.Login(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var result model.TokenResponse
	_ = json.NewDecoder(resp.Body).Decode(&result)
	assert.Equal(t, "valid-token", result.Token)
}

func TestAuthHandler_Login_InvalidPassword(t *testing.T) {
	store := &mockUserStore{users: make(map[string][]byte)}
	hashed, _ := bcrypt.GenerateFromPassword([]byte("right-password"), bcrypt.DefaultCost)
	store.users["user@example.com"] = hashed

	jwt := &mockTokenIssuer{}
	service := authservice.NewAuthService(store, jwt)
	handler := NewAuthHandler(service)

	body := model.LoginRequest{
		Email:    "user@example.com",
		Password: "wrong-password",
	}
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/api/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.Login(w, req)

	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.Contains(t, w.Body.String(), "invalid password")
}

func TestAuthHandler_Login_InvalidJSON(t *testing.T) {
	service := authservice.NewAuthService(&mockUserStore{users: make(map[string][]byte)}, &mockTokenIssuer{})
	handler := NewAuthHandler(service)

	req := httptest.NewRequest(http.MethodPost, "/api/login", bytes.NewBuffer([]byte("{")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.Login(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	assert.Contains(t, w.Body.String(), "invalid json")
}
