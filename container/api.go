package container

import (
	"github.com/crackeer/gopkg/api"
	"github.com/crackeer/gopkg/api/getter"
)

// APIRequestClient ...
var APIRequestClient *api.RequestClient

// InitAPIRequestClient ...
func InitAPIRequestClient(apiPath string) error {
	apiMetaGetter := getter.NewYamlAPIMetaGetter(apiPath)
	APIRequestClient = api.NewRequestClient(apiMetaGetter)
	return nil
}
