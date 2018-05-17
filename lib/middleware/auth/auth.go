package auth

import (
	"github.com/flexphere/lssue/lib/config"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var cnf = config.New()

func Authorize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		boards := session.Get("boards")
		if boards == nil {
			ctx.Redirect(301, "/oauth2/login")
			return
		}

		ctx.Next()
	}
}
