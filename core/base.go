package core

import (
	"database/sql"

	"github.com/stewie1520/elasticpmapi/config"
	"github.com/stewie1520/elasticpmapi/daos"
)

var _ App = (*BaseApp)(nil)

type BaseApp struct {
	config *config.Config
	dao    *daos.Dao
}

func (app *BaseApp) Config() *config.Config {
	return app.config
}

func NewApp(config *config.Config, dao *daos.Dao) *BaseApp {
	return &BaseApp{
		config: config,
		dao:    dao,
	}
}

func (app *BaseApp) Dao() *daos.Dao {
	return app.dao
}

func (app *BaseApp) DB() *sql.DB {
	if app.Dao() == nil {
		return nil
	}

	return app.Dao().DB()
}
