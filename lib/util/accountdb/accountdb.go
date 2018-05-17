package accountdb

import (
	"fmt"

	"github.com/flexphere/lssue/lib/db"
)

var QUERIES = []string{
	`CREATE TABLE pipe (id int(11) unsigned NOT NULL AUTO_INCREMENT, title VARCHAR(255) NOT NULL, created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,deleted_at timestamp NULL DEFAULT NULL, PRIMARY KEY (id)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`,
	`CREATE TABLE issue (id int(11) unsigned NOT NULL AUTO_INCREMENT, issue_id int(11) unsigned NOT NULL UNIQUE, repo VARCHAR(255) NOT NULL, title VARCHAR(255) NOT NULL, state VARCHAR(255) NOT NULL, url VARCHAR(255) NOT NULL, assignees VARCHAR(255) NOT NULL, original TEXT NOT NULL, created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,deleted_at timestamp NULL DEFAULT NULL, PRIMARY KEY (id)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`,
	`CREATE TABLE label (id int(11) unsigned NOT NULL AUTO_INCREMENT, title VARCHAR(255) NOT NULL, color VARCHAR(255) NOT NULL, bgcolor VARCHAR(255) NOT NULL, created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,deleted_at timestamp NULL DEFAULT NULL, PRIMARY KEY (id)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`,
	`CREATE TABLE category (id int(11) unsigned NOT NULL AUTO_INCREMENT, title VARCHAR(255) NOT NULL, created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,deleted_at timestamp NULL DEFAULT NULL, PRIMARY KEY (id)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`,
	`CREATE TABLE ticket (id int(11) unsigned NOT NULL AUTO_INCREMENT, title VARCHAR(255) NOT NULL, due VARCHAR(255) NULL DEFAULT "", position INT(11) NOT NULL DEFAULT 0, memo VARCHAR(255) NULL DEFAULT "", created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,deleted_at timestamp NULL DEFAULT NULL,pipe_id int(11) unsigned NOT NULL,category_id int(11) unsigned NOT NULL,PRIMARY KEY (id)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`,
	`CREATE TABLE ticket_labels (ticket_id int(11) unsigned NOT NULL, label_id int(11) unsigned NOT NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`,
	`CREATE TABLE ticket_issues (ticket_id int(11) unsigned NOT NULL, issue_id int(11) unsigned NOT NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`,
	`INSERT INTO pipe (title) VALUES ('Backlog');`,
	`INSERT INTO pipe (title) VALUES ('Icebox');`,
	`INSERT INTO pipe (title) VALUES ('Todo');`,
	`INSERT INTO pipe (title) VALUES ('inReview');`,
	`INSERT INTO pipe (title) VALUES ('Done');`,
	`INSERT INTO category (title) VALUES ('Default');`,
}

func Init(dbName string) error {
	dbName = "lssue_" + dbName

	query := fmt.Sprintf("CREATE DATABASE %s", dbName)
	if _, err := db.DB.Exec(query); err != nil {
		return err
	}

	query = fmt.Sprintf("USE %s", dbName)
	if _, err := db.DB.Exec(query); err != nil {
		return err
	}

	for _, q := range QUERIES {
		if _, err := db.DB.Exec(q); err != nil {
			return err
		}
	}

	return nil
}
