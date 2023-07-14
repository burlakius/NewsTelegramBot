package mariadb

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var botStorage *botDB

func MariadbConnect(user, password, host, dbname string) error {
	connetion, err := sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, host, dbname),
	)

	botStorage = &botDB{
		mainDB: connetion,
		mu:     sync.Mutex{},
	}

	botStorage.newsTypes = getNewsTypesFromDB()
	botStorage.adminsChats = getAdminsChatsFromDB()
	return err
}

func MariadbClose() {
	botStorage.mainDB.Close()
}

// Returns admin chats ID slice from database
func getAdminsChatsFromDB() []int64 {
	rows, err := botStorage.mainDB.Query("SELECT chat_id FROM chat_admins")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	result := make([]int64, 0, 5)
	var chat_id int64
	for rows.Next() {
		err = rows.Scan(&chat_id)
		if err != nil {
			panic(err)
		}

		result = append(result, chat_id)
	}

	return result
}

// Returns slice of news types from database
func getNewsTypesFromDB() map[string]int {
	rows, err := botStorage.mainDB.Query("SELECT name, news_type_id FROM news_types")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	result := make(map[string]int)
	var (
		newsTypeName string
		newsTypeID   int
	)
	for rows.Next() {
		err = rows.Scan(&newsTypeName, &newsTypeID)
		if err != nil {
			panic(err)
		}

		result[newsTypeName] = newsTypeID
	}

	return result
}
