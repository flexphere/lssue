package ticket

import (
	"log"
	"net/http"

	"github.com/flexphere/lssue/lib/db"
	"github.com/flexphere/lssue/lib/model"
	"github.com/flexphere/lssue/lib/util"
	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context) {
	var params model.TicketCreateParam

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

	if _, err := db.DB.NamedExec("INSERT INTO ticket (title, due, memo, pipe_id, category_id) VALUES (:title, :due, :memo, :pipe_id, :category_id);", params); err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	for _, v := range params.LabelIDS {
		db.DB.MustExec("INSERT INTO ticket_labels (ticket_id, label_id) VALUES (LAST_INSERT_ID(), ?)", v)
	}

	ctx.AbortWithStatus(http.StatusNoContent)
}
