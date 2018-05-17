package db

import (
	"fmt"

	"github.com/flexphere/lssue/lib/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	DB *sqlx.DB
)

func ConnectDB(cnf *config.Config) {
	var err error
	DB, err = sqlx.Connect(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/?parseTime=true",
			cnf.DB.User,
			cnf.DB.Pass,
			cnf.DB.Host,
			cnf.DB.Port,
		),
	)

	if err != nil {
		panic(err)
	}
}
