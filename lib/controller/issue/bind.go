package issue

import (
	"log"
	"net/http"

	"github.com/flexphere/lssue/lib/db"
	"github.com/flexphere/lssue/lib/model"
	"github.com/flexphere/lssue/lib/util"
	"github.com/gin-gonic/gin"
)

func Bind(ctx *gin.Context) {
	var params model.IssueBindParam

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

	if _, err := db.DB.NamedExec("INSERT INTO ticket_issues (ticket_id, issue_id) VALUES (:ticket_id, :id);", params); err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.AbortWithStatus(http.StatusNoContent)
}
