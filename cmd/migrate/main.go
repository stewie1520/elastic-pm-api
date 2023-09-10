//go:build migrate

package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stewie1520/elasticpmapi/config"
)

func main() {
	cfg, err := config.Init()
	handleError(err)

	migrationFilePath, ok := os.LookupEnv("MIGRATION_DIR_PATH")
	if !ok || migrationFilePath == "" {
		log.Fatalf("MIGRATION_DIR_PATH is not set")
	}

	databaseURL := cfg.DATABASE_URL

	if !strings.HasSuffix(databaseURL, "?sslmode=disable") {
		databaseURL += "?sslmode=disable"
	}

	cwd, err := os.Getwd()
	handleError(err)

	fullDir := path.Join(cwd, migrationFilePath)
	m, err := migrate.New(fmt.Sprintf("file://%s", fullDir), databaseURL)
	handleError(err)

	err = m.Up()
	defer m.Close()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Migrate: up error: %s", err)
	}

	if errors.Is(err, migrate.ErrNoChange) {
		log.Println("Migrate: no change")
		return
	}

	log.Println("Migrate: up success")
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
