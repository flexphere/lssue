package ticket

import (
	"log"
	"net/http"

	"github.com/flexphere/lssue/lib/db"
	"github.com/flexphere/lssue/lib/model"
	"github.com/flexphere/lssue/lib/util"
	"github.com/gin-gonic/gin"
)

func Sort(ctx *gin.Context) {
	var params model.TicketSortParam

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

	for i, v := range params.TicketIDS {
		if _, err := db.DB.Exec("UPDATE ticket SET pipe_id=?, position=? WHERE id=?", params.PipeID, i, v); err != nil {
			log.Println(err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	ctx.AbortWithStatus(http.StatusNoContent)
}
