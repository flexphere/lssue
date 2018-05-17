package board

import (
	"net/http"
	"strings"

	"github.com/flexphere/lssue/lib/config"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var cnf = config.New()

func List(ctx *gin.Context) {
	session := sessions.Default(ctx)
	boards := session.Get("boards")
	boardsArr := strings.Split(boards.(string), "_")
	ctx.JSON(http.StatusOK, boardsArr)
}
