package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func SuperToken(c *gin.Context) {
	supertokens.Middleware(http.HandlerFunc(
		func(rw http.ResponseWriter, r *http.Request) {
			c.Next()
		})).ServeHTTP(c.Writer, c.Request)
	c.Abort()
}
