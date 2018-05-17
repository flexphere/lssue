package ticket

import (
	"log"
	"net/http"

	"github.com/flexphere/lssue/lib/db"
	"github.com/flexphere/lssue/lib/model"
	"github.com/flexphere/lssue/lib/util"
	"github.com/gin-gonic/gin"
)

func List(ctx *gin.Context) {
	var params model.TicketListParam

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

	tickets := []model.Ticket{}
	if err := db.DB.Select(&tickets, "SELECT * FROM ticket WHERE deleted_at IS NULL ORDER BY position ASC, id DESC"); err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	for i, ticket := range tickets {
		if err := db.DB.Select(&tickets[i].LabelIDS, "SELECT label_id FROM ticket_labels WHERE ticket_id = ?", ticket.ID); err != nil {
			log.Println(err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if err := db.DB.Select(&tickets[i].IssueIDS, "SELECT issue_id FROM ticket_issues WHERE ticket_id = ?", ticket.ID); err != nil {
			log.Println(err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	ctx.JSON(http.StatusOK, tickets)
}
