package webhook

import (
	"log"
	"net/http"
	"strings"

	"github.com/flexphere/lssue/lib/db"
	"github.com/flexphere/lssue/lib/model"
	"github.com/gin-gonic/gin"
)

func Issue(ctx *gin.Context) {
	eventType := ctx.Request.Header.Get("X-GitHub-Event")

	if eventType == "ping" {
		ctx.AbortWithStatus(http.StatusOK)
	}

	if eventType != "issues" {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var params model.WebhookIssueParam
	if err := ctx.ShouldBindJSON(&params); err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	boardName := strings.ToLower(params.Repository.Owner.Login)
	if _, err := db.DB.Exec("USE lssue_" + boardName + ";"); err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	issue := &model.Issue{
		IssueID:   params.Issue.ID,
		Repo:      params.Repository.Name,
		Title:     params.Issue.Title,
		State:     params.Issue.State,
		URL:       params.Issue.HTMLURL,
		Assignees: "",
		Original:  "",
	}

	query := `
		INSERT INTO issue
			(issue_id, repo, title, state, url, assignees, original) 
		VALUES 
			(:issue_id, :repo, :title, :state, :url, :assignees, :original) 
		ON DUPLICATE KEY UPDATE 
			title=:title, state=:state, assignees=:assignees, original=:original;
	`

	if _, err := db.DB.NamedExec(query, issue); err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.AbortWithStatus(http.StatusNoContent)
}
