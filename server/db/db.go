package main

import (
	"database/sql"
	"errors"
	"fmt"
	"server/internal/config"
	"time"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	cfg := config.GetConfig()

	db, err := DoWithAttempts(func() (*sql.DB, error) {
		return sql.Open(cfg.DBDriver, cfg.DBSource)
	}, 5, 5*time.Second)

	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

func DoWithAttempts(fn func() (*sql.DB, error), attempts int, timeout time.Duration) (*sql.DB, error) {
	var err error
	var db *sql.DB

	for i := 0; i < attempts; i++ {
		db, err = fn()
		if err == nil {
			return db, nil
		}
		time.Sleep(timeout)
	}

	return nil, errors.New(fmt.Sprintf("failed after %d attempts, last error: %s", attempts, err.Error()))
}
