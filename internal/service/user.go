package service

import (
	"crypto/sha1"
	"fmt"
	"practice/internal/domain"
	"practice/internal/repository"
	"time"
)

const (
	salt = "someSalt"
	//signingKey = "sfesf25r23efer23rf"
	//tokenTTL   = 12 * time.Hour
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) SignUp(input domain.SignUpInput) (int, error) {
	password := generatePasswordHash(input.Password)

	user := domain.User{
		Name:         input.Name,
		Email:        input.Email,
		Password:     password,
		RegisteredAt: time.Now(),
	}

	return s.repo.SignUp(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
