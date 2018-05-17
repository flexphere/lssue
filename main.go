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
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	melody "gopkg.in/olahol/melody.v1"
)

func main() {
	config := config.New()

	err := db.ConnectDB(config)
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.Use(gin.Recovery())

	store := cookie.NewStore([]byte(config.Session.Secret))
	r.Use(sessions.Sessions(config.Session.Name, store))

	r.GET("/ping", common.Healthcheck)
	r.GET("/favicon.ico", common.Favicon)
	r.GET("/oauth2/login", oauth2.Login)
	r.GET("/oauth2/callback", oauth2.Callback)
	r.POST("/webhook/issue", webhook.Issue)

	r.Use(auth.Authorize())

	r.Use(static.Serve("/", static.LocalFile("./static/dist", true)))

	api := r.Group("/api")
	{
		api.POST("/board", board.List)
		api.POST("/pipe", pipe.List)
		api.POST("/ticket", ticket.List)
		api.POST("/ticket/create", ticket.Create)
		api.POST("/ticket/update", ticket.Update)
		api.POST("/ticket/delete", ticket.Delete)
		api.POST("/ticket/sort", ticket.Sort)
		api.POST("/category", category.List)
		api.POST("/category/create", category.Create)
		api.POST("/category/delete", category.Delete)
		api.POST("/label", label.List)
		api.POST("/label/create", label.Create)
		api.POST("/label/delete", label.Delete)
		api.POST("/issue", issue.List)
		api.POST("/issue/bind", issue.Bind)
		api.POST("/issue/unbind", issue.Unbind)
	}

	// Websockets
	m := melody.New()
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.BroadcastFilter(msg, func(q *melody.Session) bool {
			return q.Request.URL.Path == s.Request.URL.Path
		})
	})

	r.GET("/ws/:name", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	r.Run()
}
