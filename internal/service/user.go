package service

import (
	"context"
	"practice/internal/domain"
	"time"
)

type PasswordHasher interface {
	Hash(password string) (string, error)
}

type UsersRepository interface {
}

type Users struct {
	repo   UsersRepository
	hasher PasswordHasher

	hmacSecret []byte
}

func NewUsers(repo UsersRepository, hasher PasswordHasher, secret []byte) *Users {
	return &Users{
		repo:       repo,
		hasher:     hasher,
		hmacSecret: secret,
	}
}

func (s *Users) SignUp(ctx context.Context, input domain.SignUpInput) error {
	password, err := s.hasher.Hash(input.Password)
	if err != nil {
		return err
	}

	user := domain.User{
		Name:         input.Name,
		Email:        input.Email,
		Password:     password,
		RegisteredAt: time.Now(),
	}

	return s.repo
}
