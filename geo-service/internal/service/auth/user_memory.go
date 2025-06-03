package auth

import (
	"errors"
	"sync"
)

type UserStore interface {
	Save(email string, hashedPassword []byte) error
	Get(email string) ([]byte, error)
}

type MemoryUserStore struct {
	mu    sync.RWMutex
	users map[string][]byte
}

func NewMemoryUserStore() *MemoryUserStore {
	return &MemoryUserStore{
		users: make(map[string][]byte),
	}
}

func (s *MemoryUserStore) Save(email string, hashedPassword []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.users[email]; ok {
		return errors.New("user with this email already exists")
	}
	s.users[email] = hashedPassword
	return nil
}

func (s *MemoryUserStore) Get(email string) ([]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	hashedPassword, ok := s.users[email]
	if !ok {
		return nil, errors.New("user not found")
	}

	return hashedPassword, nil
}
