package core

import (
	"database/sql"

	"github.com/stewie1520/elasticpmapi/config"
	"github.com/stewie1520/elasticpmapi/daos"
)

type App interface {
	DB() *sql.DB
	Dao() *daos.Dao

	Bootstrap() error
	Config() *config.Config
	IsDebug() bool
}
