package container

import (
	"database/sql"
	"sync"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	dbLocker *sync.Mutex
	DB       *gorm.DB
)

func InitDB() {
	dbLocker = &sync.Mutex{}
	db, err := gorm.Open(sqlite.Open(config.Sqlite3DatabaseFile), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
}

func GetDatabase() *gorm.DB {
	return DB
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
