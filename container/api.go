package container

import (
	"github.com/crackeer/gopkg/api"
	"github.com/crackeer/gopkg/api/getter"
)

var APIRequestClient *api.RequestClient

// InitAPIRequestClient
func InitAPIRequestClient() error {
	apiMetaGetter := getter.NewYamlAPIMetaGetter("config/api")
	APIRequestClient = api.NewRequestClient(apiMetaGetter)
	return nil
}
