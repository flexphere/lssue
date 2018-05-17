package issue

import (
	"log"
	"net/http"

	"github.com/flexphere/lssue/lib/db"
	"github.com/flexphere/lssue/lib/model"
	"github.com/flexphere/lssue/lib/util"
	"github.com/gin-gonic/gin"
)

func List(ctx *gin.Context) {
	var params model.IssueListParam

	if err := ctx.ShouldBindJSON(&params); err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := util.VerifyBoard(ctx, params.Boardname); err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	pipes := []model.Issue{}
	if err := db.DB.Select(&pipes, "SELECT * FROM issue WHERE deleted_at IS NULL ORDER BY id ASC"); err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, pipes)
}
