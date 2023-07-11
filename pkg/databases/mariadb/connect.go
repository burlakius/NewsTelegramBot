package mariadb

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

func MariadbConnect(user, password, host, dbname string) error {
	connetion, err := sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, host, dbname),
	)

	botStorage = &botDB{
		mainDB: connetion,
		mu:     sync.Mutex{},
	}

	botStorage.adminsChats = getAdminsChatsFromDB()
	return err
}

func MariadbClose() {
	botStorage.mainDB.Close()
}
