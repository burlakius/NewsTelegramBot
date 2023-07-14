package mariadb

import (
	"database/sql"
	"sync"
)

type botDB struct {
	mu          sync.Mutex
	mainDB      *sql.DB
	adminsChats []int64
	newsTypes   map[string]int
}

type News struct {
	ChatID    int64
	MessageID int
}

type Question struct {
	UserID    int64
	FirstName string
	LastName  string
	Username  string
	ChatID    int64
	MessageID int
}
