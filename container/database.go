package container

import (
	"fmt"
	"strings"

	"github.com/crackeer/gopkg/config"
	"github.com/crackeer/gopkg/storage"
	"github.com/crackeer/gopkg/util"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/gorm"
)

var (

	// DBMap
	ReadDB   map[string]*gorm.DB
	WriteDB  map[string]*gorm.DB
	DBDriver map[string]string
)

// InitDB
//  @param sqliteDBFile
//  @param boltDBFile
func InitDB(dir string, env string) error {

	ReadDB = map[string]*gorm.DB{}
	WriteDB = map[string]*gorm.DB{}
	DBDriver = map[string]string{}
	fileList := util.GetFiles(dir)
	fmt.Println(dir, env, fileList)
	for _, file := range fileList {
		if !strings.HasSuffix(file, config.YamlExt) {
			continue
		}
		conf, err := loadDBConfig(file, env)
		if err != nil {
			return err
		}
		name := lastFileName(file)
		readDB, writeDB, err := openDB(conf, name)
		if err != nil {
			return err
		}
		WriteDB[name] = writeDB
		ReadDB[name] = readDB
		DBDriver[name] = conf.Driver

	}
	return nil
}

func lastFileName(path string) string {
	parts := strings.Split(path, "/")
	lastItem := parts[len(parts)-1]
	return strings.TrimRight(lastItem, config.YamlExt)
}

func loadDBConfig(file string, env string) (*config.DBConfig, error) {
	mapConfig := map[string]*config.DBConfig{}
	if err := config.LoadYamlConfig(file, &mapConfig); err != nil {
		return nil, fmt.Errorf("load database config `%s` error: %s", file, err.Error())
	}

	if conf, exists := mapConfig[env]; exists {
		return conf, nil
	}

	return nil, fmt.Errorf("file `%s` env `%s` database config not found", file, env)
}

func openDB(conf *config.DBConfig, name string) (*gorm.DB, *gorm.DB, error) {

	var (
		readDB, writeDB *gorm.DB
		err             error
	)
	if conf.Driver == storage.DriverMySQL {
		writeDB, err = storage.GetMySQLDB(&storage.MySQLConfig{
			User:     conf.WriteUser,
			Password: conf.WritePassword,
			Host:     conf.WriteHost,
			Charset:  conf.Charset,
			Database: conf.Database,
		}, nil)

		if err != nil {
			return nil, nil, fmt.Errorf("open write db `%s` error: %s", name, err.Error())
		}
		readDB, err = storage.GetMySQLDB(&storage.MySQLConfig{
			User:     conf.ReadUser,
			Password: conf.ReadPassword,
			Host:     conf.ReadHost,
			Charset:  conf.Charset,
			Database: conf.Database,
		}, nil)
		if err != nil {
			return nil, nil, fmt.Errorf("open read db `%s` error: %s", name, err.Error())
		}
	} else if conf.Driver == storage.DriverSQLite {
		writeDB, err = storage.GetSQliteDB(conf.File, nil)
		if err != nil {
			return nil, nil, fmt.Errorf("open write db `%s` error: %s", name, err.Error())
		}
		readDB = writeDB
	}

	return readDB, writeDB, nil
}

// GetWriteDB
//  @param name
//  @return *gorm.DB
func GetWriteDB(name string) *gorm.DB {
	fmt.Println("GetWriteDB", name, WriteDB[name])
	if db, exists := WriteDB[name]; exists {
		return db
	}
	return nil
}

// GetReadDB
//  @param name
//  @return *gorm.DB
func GetReadDB(name string) *gorm.DB {
	if db, exists := ReadDB[name]; exists {
		return db
	}
	return nil
}

func GetDatabase() *gorm.DB {
	return nil
}
