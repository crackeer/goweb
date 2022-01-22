package container

import (
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitDB() {
	db, err := gorm.Open(sqlite.Open(config.Sqlite3DatabaseFile), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
}

func GetDatabase() *gorm.DB {
	return DB
}
