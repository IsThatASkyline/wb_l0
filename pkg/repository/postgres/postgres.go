package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
)

func NewPostgresDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=postgresql port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD")))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
