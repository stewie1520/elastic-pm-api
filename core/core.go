package core

import "github.com/stewie1520/elasticpmapi/config"

type App interface {
	Config() *config.Config
}
