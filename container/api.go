package container

import (
	"github.com/crackeer/gopkg/api"
	"github.com/crackeer/gopkg/api/getter"
)

var APIRequestClient *api.RequestClient

// InitAPIRequestClient
func InitAPIRequestClient() {
	apiMetaGetter := getter.NewYamlAPIMetaGetter("config/api")
	APIRequestClient = api.NewRequestClient(apiMetaGetter)
}
