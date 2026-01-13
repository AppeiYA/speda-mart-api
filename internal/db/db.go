package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	*sqlx.DB
}

func ConnectDB(DatabaseUrl string) (*DB, error) {
	db, err := sqlx.Connect("postgres", DatabaseUrl)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to db: %w", err)
	}

	if pingerr := db.Ping(); pingerr != nil {
		return nil, fmt.Errorf("Failed to ping db: %w", pingerr)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	return &DB{db}, nil
}