package db

import (
	"fmt"
	"log"

	"github.com/flexphere/lssue/lib/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	DB *sqlx.DB
)

func ConnectDB(cnf *config.Config) error {
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
	return err
}

func DatabaseExists(database string) bool {
	_, err := DB.Exec("SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = '$1'", database)
	if err != nil {
		log.Panic(err)
		return false
	}
	return true
}
