package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func NewPostgresDB(connectionURL string) (*sql.DB, error) {
	conn, err := sql.Open("postgres", connectionURL)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}

	fmt.Println("Connected to database 🎉")

	return conn, nil
}
