package auth

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"sync"
)

type AuthService struct {
	users UserStore
	mu    sync.RWMutex
	jwt   TokenIssuer
}

func NewAuthService(users UserStore, jwt TokenIssuer) *AuthService {
	return &AuthService{
		users: users,
		jwt:   jwt,
	}
}

func (s *AuthService) Register(email, password string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, err := s.users.Get(email); err == nil {
		return errors.New("user with email already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	return s.users.Save(email, hash)
}

func (s *AuthService) Login(email, password string) (string, error) {
	s.mu.RLock()
	hashedPass, err := s.users.Get(email)
	s.mu.RUnlock()

	if err != nil {
		return "", errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword(hashedPass, []byte(password)); err != nil {
		return "", errors.New("invalid password")
	}

	return s.jwt.Issue(email)
}
