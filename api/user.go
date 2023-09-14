package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stewie1520/elasticpmapi/api/middleware"
	"github.com/stewie1520/elasticpmapi/core"
	"github.com/stewie1520/elasticpmapi/usecases"
	"github.com/supertokens/supertokens-golang/recipe/session"
)

type userApi struct {
	app core.App
}

func bindUserApi(app core.App, ginEngine *gin.Engine) *userApi {
	api := &userApi{
		app: app,
	}

	subGroup := ginEngine.Group("/user")
	subGroup.GET("/me", middleware.VerifySession(nil), api.getUser)

	return api
}

func (api *userApi) getUser(c *gin.Context) {
	q := usecases.NewGetUserByAccountIDQuery(api.app)
	sessionContainer := session.GetSessionFromRequestContext(c.Request.Context())
	// TODO: this is wrong, fix later
	q.AccountID = sessionContainer.GetUserID()

	if user, err := q.Execute(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, user)
	}

}
