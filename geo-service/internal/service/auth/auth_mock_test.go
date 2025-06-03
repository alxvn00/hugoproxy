package auth

import (
	"errors"
	"sync"
)

type MockUserStore struct {
	mu   sync.RWMutex
	data map[string][]byte
}

func NewMockUserStore() *MockUserStore {
	return &MockUserStore{data: make(map[string][]byte)}
}

func (m *MockUserStore) Get(email string) ([]byte, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	val, ok := m.data[email]
	if !ok {
		return nil, errors.New("user not found")
	}
	return val, nil
}

func (m *MockUserStore) Save(email string, hash []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.data[email]; exists {
		return errors.New("already exists")
	}
	m.data[email] = hash
	return nil
}

type MockTokenIssuer struct {
	token string
	err   error
}

func (m *MockTokenIssuer) Issue(email string) (string, error) {
	return m.token, m.err
}
