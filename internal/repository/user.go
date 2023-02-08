package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"practice/internal/config"
	"practice/internal/domain"
)

type User struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) *User {
	return &User{db: db}
}

func (u *User) SignUp(user domain.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, email, password, registered_at) values ($1, $2, $3, $4) RETURNING id", config.UsersTable)
	row := u.db.QueryRow(query, user.Name, user.Email, user.Password, user.RegisteredAt)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
