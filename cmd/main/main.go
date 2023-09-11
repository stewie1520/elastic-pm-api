package main

import (
	"fmt"

	"github.com/stewie1520/elasticpmapi/api"
	"github.com/stewie1520/elasticpmapi/auth"
	"github.com/stewie1520/elasticpmapi/config"
	"github.com/stewie1520/elasticpmapi/core"
	"github.com/stewie1520/elasticpmapi/daos"
	"github.com/stewie1520/elasticpmapi/db"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		panic(err.Error())
	}

	if err := auth.InitSuperToken(cfg); err != nil {
		panic(err.Error())
	}

	db, err := db.NewPostgresDB(cfg)
	if err != nil {
		panic(err.Error())
	}

	dao := daos.New(db)
	app := core.NewApp(cfg, dao)

	router, err := api.InitApi(app)
	if err != nil {
		panic(err.Error())
	}

	router.Run(fmt.Sprintf(":%d", cfg.Port))
}
