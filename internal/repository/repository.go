package repository

import (
	"github.com/jmoiron/sqlx"
	"practice/internal/domain"
)

type Authorization interface {
	SignUp(user domain.User) (int, error)
}

type Repository struct {
	Authorization
}

func NewRepositories(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewUser(db),
	}
}
