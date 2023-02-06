package repository

import (
	"context"
	"database/sql"
	"practice/internal/domain"
)

type Users struct {
	db *sql.DB
}

func NewUsers(db *sql.DB) *Users {
	return &Users{db}
}

func (r *Users) Create(ctx context.Context, user domain.User) error {
	_, err := r.db.Exec("INSERT INTO users (name, email, passowrd, registered_at) values ($1, $2, $3, $4)",
		user.Name, user.Email, user.Password, user.RegisteredAt)

	return err
}
