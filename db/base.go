package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/stewie1520/elasticpmapi/config"
)

func NewPostgresDB(cfg *config.Config) (*sql.DB, error) {
	conn, err := sql.Open("postgres", cfg.DATABASE_URL)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}

	fmt.Println("Connected to database 🎉")

	return conn, nil
}
