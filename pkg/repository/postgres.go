package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type DBConfig struct {
	Host string
	Port string
	User string
	Pass string
	Name string
	SSL  string
}

func NewPostgresDB(cfg DBConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Name, cfg.Pass, cfg.SSL))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, err
}
