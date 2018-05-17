package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Healthcheck(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusOK)
}

func Favicon(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusNotFound)
}
