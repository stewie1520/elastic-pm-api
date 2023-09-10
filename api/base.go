package api

import (
	"github.com/gin-gonic/gin"
	"github.com/stewie1520/elasticpmapi/api/middleware"
	"github.com/stewie1520/elasticpmapi/core"
)

func InitApi(app core.App) (*gin.Engine, error) {
	engine := gin.New()
	engine.Use(middleware.Cors(app))
	engine.Use(middleware.SuperToken)

	return engine, nil
}
