package server

import (
	"fmt"
	"net/http"

	"github.com/crackeer/goweb/define"
	"github.com/crackeer/goweb/server/handler"
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
	//下载数据库文件
	router.StaticFile("download/sqlite", config.Resource.SqliteDBFile)
	router.StaticFile("download/bolt", config.Resource.BoltDBFile)

	setupAPIRouter(router)

	router.NoRoute(pager.Render)
	return router.Run(fmt.Sprintf(":%d", config.Port))
}

func setupAPIRouter(router *gin.Engine) {
	apiRouter := router.Group("api/v1")
	apiRouter.GET("list/:table", table.GetList)

	apiRouter.POST("object/update", handler.UpdateObject)
	apiRouter.POST("object/upload", handler.UploadObject)
	apiRouter.POST("object/delete", handler.DeleteObject)
	apiRouter.POST("object/share", handler.ShareObject)
	apiRouter.GET("test", test.RequestAPI)
}
