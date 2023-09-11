package main

import (
	"flag"
	"fmt"

	"github.com/stewie1520/elasticpmapi/api"
	"github.com/stewie1520/elasticpmapi/config"
	"github.com/stewie1520/elasticpmapi/core"
	"github.com/stewie1520/elasticpmapi/usecases"
)

var debug = flag.Bool("debug", false, "debug mode")

func init() {
	flag.Parse()
}

func main() {
	cfg, err := config.Init()
	panicIfError(err)

	app := core.NewBaseApp(core.BaseAppConfig{
		Config:  cfg,
		IsDebug: *debug,
	})

	err = app.Bootstrap()
	panicIfError(err)

	usecases.AddHandlersToHook(app)

	router, err := api.InitApi(app)
	panicIfError(err)

	router.Run(fmt.Sprintf(":%d", cfg.Port))
}

func panicIfError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
