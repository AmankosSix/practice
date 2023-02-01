package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"practice/internal/config"
)

func NewClient(cfg config.Postgres) (*sqlx.DB, error) {
	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Name, cfg.Password, cfg.SSLMode)

	db, err := sqlx.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
