package util

import (
	"errors"
	"strings"

	"github.com/flexphere/lssue/lib/config"
	"github.com/flexphere/lssue/lib/db"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var cnf = config.New()

func VerifyBoard(ctx *gin.Context, boardName string) error {
	session := sessions.Default(ctx)
	boards := session.Get("boards")

	if exists := strings.Contains(boards.(string), boardName); !exists {
		return errors.New("board_name doesn't exist in cookie")
	}

	if _, err := db.DB.Exec("use lssue_" + boardName + ";"); err != nil {
		return errors.New("error while switching database")
	}

	return nil
}
