package container

import (
	"github.com/crackeer/gopkg/config"
	"github.com/crackeer/goweb/define"
)

var AppConfig *define.AppConfig

// InitConfig
//  @param configPath
//  @return error
func InitConfig(configPath string) error {
	AppConfig = &define.AppConfig{}
	return config.LoadYamlConfig(configPath, AppConfig)
}
