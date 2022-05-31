package container

import (
	"github.com/boltdb/bolt"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	boltDB *bolt.DB
)

func InitDB() {
	db, err := gorm.Open(sqlite.Open(config.Database), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
	initBoltDB()
}

func GetDatabase() *gorm.DB {
	return DB
}

func getBoltDatabase() *bolt.DB {
	return boltDB
}

func initBoltDB() error {
	db, err := bolt.Open(config.BoltDB, 0600, nil)
	if err != nil {
		return err
	}
	boltDB = db
	return nil
	/*
		defer db.Close()
		updateErr := db.Update(func(tx *bolt.Tx) error {
			b, err := tx.CreateBucket([]byte("MyBucket"))
			if err != nil {
				return fmt.Errorf("create bucket: %s", err)
			}
			err = b.Put([]byte("answer"), []byte("42"))
			return err
		})
		fmt.Println(updateErr)
		tx, _ := db.Begin(false)
		val := tx.Bucket([]byte("MyBucket")).Get([]byte("answer"))

		bucket1 := tx.Bucket([]byte("test"))
		bucket2 := tx.Bucket([]byte("MyBucket"))
		fmt.Println(string(val), bucket1, bucket2)
	*/
}
