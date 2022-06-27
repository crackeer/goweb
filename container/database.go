package container

import (
	"errors"
	"io/ioutil"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/crackeer/gopkg/util"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (

	// DBMap 
	DBMap map[string]*gorm.DB
	boltDB   *bolt.DB
)

// InitDB
//  @param sqliteDBFile
//  @param boltDBFile
func InitDB(dir string) error {
	fileList := util.GetFiles(dir)
	for _, file := range fileList {
		if !strings.HasSuffix(file, ".yaml") {
			continue
		}

	}
}

// GetDatabase 
//  @param name 
//  @return *gorm.DB 
func GetDatabase(name string) *gorm.DB {
	if db, exists := DBMap[name];exists {
		return db
	}
	return nil
}

/*

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
*/
