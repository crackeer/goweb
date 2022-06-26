package container

import (
	"errors"

	"github.com/boltdb/bolt"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	SqliteDB *gorm.DB
	boltDB   *bolt.DB
)

// InitDB
//  @param sqliteDBFile
//  @param boltDBFile
func InitDB(sqliteDBFile string, boltDBFile string) error {
	db, err := gorm.Open(sqlite.Open(sqliteDBFile), &gorm.Config{})
	if err != nil {
		return errors.New("failed to connect database")
	}
	SqliteDB = db
	if len(boltDBFile) > 0 {
		initBoltDB(boltDBFile)
	}

	return nil
}

// GetDatabase
//  @return *gorm.DB
func GetDatabase() *gorm.DB {
	return SqliteDB
}

// GetBoltDatabase
//  @return *bolt.DB
func GetBoltDatabase() *bolt.DB {
	return boltDB
}

func initBoltDB(boltDBFile string) error {
	db, err := bolt.Open(boltDBFile, 0600, nil)
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
