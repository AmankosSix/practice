package service

import "practice/internal/repository"

type Services struct {
}

func NewServices(repos *repository.Repositories) *Services {
	return &Services{}
}
