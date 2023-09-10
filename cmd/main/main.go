package main

import (
	"fmt"

	"github.com/stewie1520/elasticpmapi/api"
	"github.com/stewie1520/elasticpmapi/auth"
	"github.com/stewie1520/elasticpmapi/config"
	"github.com/stewie1520/elasticpmapi/core"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		panic(err.Error())
	}

	if err := auth.InitSuperToken(cfg); err != nil {
		panic(err.Error())
	}

	app := core.NewApp(cfg)
	router, err := api.InitApi(app)
	if err != nil {
		panic(err.Error())
	}

	router.Run(fmt.Sprintf(":%d", cfg.Port))
}
