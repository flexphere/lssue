package ticket

import (
	"log"
	"net/http"

	"github.com/flexphere/lssue/lib/db"
	"github.com/flexphere/lssue/lib/model"
	"github.com/flexphere/lssue/lib/util"
	"github.com/gin-gonic/gin"
)

func Update(ctx *gin.Context) {
	var params model.TicketUpdateParam

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

	if _, err := db.DB.NamedExec("UPDATE ticket SET title=:title, due=:due, memo=:memo, pipe_id=pipe_id, category_id=:category_id WHERE id=:id;", params); err != nil {
		log.Panicln(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if _, err := db.DB.Exec("DELETE FROM ticket_labels WHERE ticket_id=?", params.ID); err != nil {
		log.Panicln(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	for _, v := range params.LabelIDS {
		db.DB.MustExec("INSERT INTO ticket_labels (ticket_id, label_id) VALUES (?, ?)", params.ID, v)
	}

	ctx.AbortWithStatus(http.StatusNoContent)
}
