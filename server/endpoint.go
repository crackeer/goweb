package server

import (
	"fmt"
	"net/http"

	"github.com/crackeer/gopkg/ginhelper"
	"github.com/crackeer/goweb/define"
	"github.com/crackeer/goweb/server/handler"
	"github.com/crackeer/goweb/server/handler/database"
	pager "github.com/crackeer/goweb/server/handler/page"
	"github.com/crackeer/goweb/server/handler/table"
	"github.com/crackeer/goweb/server/handler/test"
	"github.com/gin-gonic/gin"
)

func Run(config *define.AppConfig) error {

	router := gin.New()
	gin.SetMode(gin.DebugMode)

	//前端静态文件
	router.StaticFS("/public", http.Dir(config.Resource.PublicDir))
	router.GET("image/:id", handler.RenderImage)

	setupAPIRouter(router)

	router.NoRoute(pager.Render)
	return router.Run(fmt.Sprintf(":%d", config.Port))
}

func setupAPIRouter(router *gin.Engine) {
	apiRouter := router.Group("api/v1/:database", ginhelper.DoResponseJSON())
	apiRouter.GET("query/:table", table.Query)
	apiRouter.GET("list/:table", table.List)
	apiRouter.POST("create/:table", table.Create)
	apiRouter.POST("update/:table", table.Update)
	apiRouter.POST("delete/:table", table.Delete)
	apiRouter.GET("distinct/:table", table.Distinct)
	apiRouter.POST("exec/sql", database.Exec)
	apiRouter.GET("tables", database.Tables)

	apiRouter.POST("object/upload", handler.UploadObject)
	apiRouter.POST("object/share", handler.ShareObject)
	apiRouter.GET("test", test.RequestAPI)
}
