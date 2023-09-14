package db

import (
	"database/sql"
	"time"
)

type Option func(*sql.DB)

func WithMaxIdleConns(n int) Option {
	return func(db *sql.DB) {
		db.SetMaxIdleConns(n)
	}
}

func WithMaxOpenConns(n int) Option {
	return func(db *sql.DB) {
		db.SetMaxOpenConns(n)
	}
}

func WithConnMaxIdleTime(t time.Duration) Option {
	return func(db *sql.DB) {
		db.SetConnMaxIdleTime(t * time.Minute)
	}
}
