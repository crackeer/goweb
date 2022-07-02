package page

import (
	"github.com/crackeer/gopkg/config"
	"github.com/crackeer/goweb/define"
)

func loadPageConfig(path string) (*define.PageConfig, error) {
	retData := &define.PageConfig{}
	err := config.LoadYamlConfig(path, retData)
	return retData, err
}
