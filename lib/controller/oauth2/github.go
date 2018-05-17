package oauth2

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/flexphere/lssue/lib/config"
	"github.com/flexphere/lssue/lib/util/accountdb"
	"github.com/flexphere/lssue/lib/util/randstr"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	githuboauth "golang.org/x/oauth2/github"
)

var cnf = config.New()

var oauth = oauth2.Config{
	ClientID:     cnf.Github.ClientID,
	ClientSecret: cnf.Github.ClientSecret,
	Endpoint:     githuboauth.Endpoint,
	RedirectURL:  cnf.Github.RedirectURL,
	Scopes:       cnf.Github.Scopes,
}

func Login(ctx *gin.Context) {
	session := sessions.Default(ctx)

	state := randstr.String(16)
	session.Set("state", state)
	session.Save()

	url := oauth.AuthCodeURL(state, oauth2.AccessTypeOnline)
	ctx.Writer.WriteString(fmt.Sprintf(`<!DOCTYPE html><html><head><meta http-equiv="refresh" content="0;URL=%s"></head><body></body></html>`, url))
	ctx.AbortWithStatus(http.StatusOK)
}

func Callback(ctx *gin.Context) {
	session := sessions.Default(ctx)

	sessState := session.Get("state")
	if sessState == nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error Session STATE"))
		return
	}

	state := ctx.Query("state")
	if state != sessState.(string) {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error STATE"))
		return
	}

	code := ctx.Query("code")
	token, err := oauth.Exchange(oauth2.NoContext, code)
	if err != nil {
		panic(err)
	}

	oauthClient := oauth.Client(oauth2.NoContext, token)
	client := github.NewClient(oauthClient)

	boards := []string{}

	user, _, err1 := client.Users.Get(ctx, "")
	if err1 != nil {
		panic(err1)
	}
	boards = append(boards, strings.ToLower(*user.Login))

	orgs, _, err2 := client.Organizations.List(ctx, "", &github.ListOptions{})
	if err2 != nil {
		panic(err2)
	}

	for _, org := range orgs {
		boards = append(boards, strings.ToLower(*org.Login))
	}

	for _, dbname := range boards {
		if err := accountdb.Init(dbname); err != nil {
			log.Println(err)
		}
	}

	boardsStr := strings.Join(boards, "_")

	session.Set("boards", boardsStr)
	session.Save()

	ctx.Writer.WriteString(`<!DOCTYPE html><html><head><meta http-equiv="refresh" content="0;URL=/"></head><body></body></html>`)
}
