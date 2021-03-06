package server

import (
	"fmt"

	ginhelper "github.com/crackeer/gopkg/gin"
	"github.com/crackeer/goweb/define"
	"github.com/crackeer/goweb/server/handler/api"
	pager "github.com/crackeer/goweb/server/handler/page"
	"github.com/gin-gonic/gin"
)

// Run ...
//  @param config
//  @return error
func Run(config *define.AppConfig) error {

	router := gin.New()
	gin.SetMode(gin.DebugMode)
	router.Any("/api/:service/:api", ginhelper.DoResponseJSON(), api.RequestAPI)
	router.NoRoute(pager.Render)
	return router.Run(fmt.Sprintf(":%d", config.Port))
}
