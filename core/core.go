package core

import (
	"database/sql"

	"github.com/stewie1520/elasticpmapi/config"
	"github.com/stewie1520/elasticpmapi/daos"
	hook "github.com/stewie1520/elasticpmapi/hooks"
)

type App interface {
	DB() *sql.DB
	Dao() *daos.Dao

	Bootstrap() error
	Config() *config.Config
	IsDebug() bool

	// OnAfterAccountCreated hook is triggered after an account is created in identity service (SuperTokens for e.g)
	// This is useful when you want to create an user in your database after an account is created in identity service
	OnAfterAccountCreated() *hook.Hook[*AccountCreatedEvent]
}
