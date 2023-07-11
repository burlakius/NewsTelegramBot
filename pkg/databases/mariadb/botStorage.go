package mariadb

import (
	"database/sql"
	"sync"
)

var botStorage *botDB

type botDB struct {
	mu          sync.Mutex
	mainDB      *sql.DB
	adminsChats []int64
}
