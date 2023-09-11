package core

import (
	"database/sql"

	"github.com/stewie1520/elasticpmapi/config"
	"github.com/stewie1520/elasticpmapi/daos"
)

type App interface {
	Config() *config.Config

	DB() *sql.DB
	Dao() *daos.Dao
}
