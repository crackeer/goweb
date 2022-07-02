package server

import (
	"fmt"

	"github.com/crackeer/goweb/define"
	pager "github.com/crackeer/goweb/server/handler/page"
	"github.com/gin-gonic/gin"
)

// Run ...
//  @param config
//  @return error
func Run(config *define.AppConfig) error {

	router := gin.New()
	gin.SetMode(gin.DebugMode)
	router.NoRoute(pager.Render)
	return router.Run(fmt.Sprintf(":%d", config.Port))
}
