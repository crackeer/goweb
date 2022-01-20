package container

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/crackeer/goweb/define"
)

var config *define.Config

// Init
//  @param configPath
//  @return error
func Init(configPath string) error {
	bytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("read config file error, %s", err.Error())
	}

	config = &define.Config{}
	decodeError := json.Unmarshal(bytes, config)
	if decodeError != nil {
		return fmt.Errorf("unmarshal config content error, %s", decodeError.Error())
	}
	InitDB()
	LockDatabase()
	UnlockDatabase()

	return nil
}

// GetConfig
//  @return *define.Config
func GetConfig() *define.Config {
	return config
}
