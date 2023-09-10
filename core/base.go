package core

import "github.com/stewie1520/elasticpmapi/config"

var _ App = (*BaseApp)(nil)

type BaseApp struct {
	config *config.Config
}

func (app *BaseApp) Config() *config.Config {
	return app.config
}

func NewApp(config *config.Config) *BaseApp {
	return &BaseApp{
		config: config,
	}
}
