package ws

import (
	"github.com/gin-gonic/gin"
	melody "gopkg.in/olahol/melody.v1"
)

var (
	Melody *melody.Melody
)

func Init() {
	Melody = melody.New()
	Melody.HandleMessage(func(s *melody.Session, msg []byte) {
		Melody.BroadcastFilter(msg, func(q *melody.Session) bool {
			return q.Request.URL.Path == s.Request.URL.Path
		})
	})
}

func Handler(c *gin.Context) {
	Melody.HandleRequest(c.Writer, c.Request)
}
