package mariadb

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var db *botDB

type botDB struct {
	cursor *sql.DB
	mu     sync.Mutex
}

func MariadbConnect(user, password, host, dbname string) error {
	connetion, err := sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, host, dbname),
	)

	db = &botDB{
		cursor: connetion,
	}

	return err
}
