package service

import (
	"practice/internal/domain"
	"practice/internal/repository"
)

type Authorization interface {
	SignUp(user domain.SignUpInput) (int, error)
}

type Services struct {
	Authorization
}

func NewServices(repos *repository.Repository) *Services {
	return &Services{
		Authorization: NewAuthService(repos.Authorization),
	}
}
