package main

import (
	"github.com/flexphere/lssue/lib/config"
	"github.com/flexphere/lssue/lib/controller/board"
	"github.com/flexphere/lssue/lib/controller/category"
	"github.com/flexphere/lssue/lib/controller/common"
	"github.com/flexphere/lssue/lib/controller/issue"
	"github.com/flexphere/lssue/lib/controller/label"
	"github.com/flexphere/lssue/lib/controller/oauth2"
	"github.com/flexphere/lssue/lib/controller/pipe"
	"github.com/flexphere/lssue/lib/controller/ticket"
	"github.com/flexphere/lssue/lib/controller/webhook"
	"github.com/flexphere/lssue/lib/db"
	"github.com/flexphere/lssue/lib/middleware/auth"
	"github.com/flexphere/lssue/lib/ws"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	config := config.New()
	store := cookie.NewStore([]byte(config.Session.Secret))

	ws.Init()
	db.ConnectDB(config)

	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(sessions.Sessions(config.Session.Name, store))

	r.GET("/ping", common.Healthcheck)
	r.GET("/favicon.ico", common.Favicon)
	r.GET("/oauth2/login", oauth2.Login)
	r.GET("/oauth2/callback", oauth2.Callback)
	r.POST("/webhook/issue", webhook.Issue)

	r.Use(auth.Authorize())
	r.Use(static.Serve("/", static.LocalFile("./static/dist", true)))

	r.GET("/ws/:name", ws.Handler)
	r.POST("/api/board", board.List)
	r.POST("/api/pipe", pipe.List)
	r.POST("/api/ticket", ticket.List)
	r.POST("/api/ticket/create", ticket.Create)
	r.POST("/api/ticket/update", ticket.Update)
	r.POST("/api/ticket/delete", ticket.Delete)
	r.POST("/api/ticket/sort", ticket.Sort)
	r.POST("/api/category", category.List)
	r.POST("/api/category/create", category.Create)
	r.POST("/api/category/delete", category.Delete)
	r.POST("/api/label", label.List)
	r.POST("/api/label/create", label.Create)
	r.POST("/api/label/delete", label.Delete)
	r.POST("/api/issue", issue.List)
	r.POST("/api/issue/bind", issue.Bind)
	r.POST("/api/issue/unbind", issue.Unbind)

	r.Run()
}
