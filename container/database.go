package container

import (
	"database/sql"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var (
	dbLocker *sync.Mutex
)

func init() {
	dbLocker = &sync.Mutex{}
}

// LockDatabase
//  @return *sql.DB
//  @return error
func LockDatabase() (*sql.DB, error) {
	dbLocker.Lock()
	return sql.Open("sqlite3", config.Sqlite3DatabaseFile)
}

// UnlockDatabase
func UnlockDatabase() {
	dbLocker.Unlock()
}
