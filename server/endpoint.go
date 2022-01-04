package server

import (
	"fmt"
	"net/http"

	"github.com/crackeer/goweb/container"
	"github.com/crackeer/goweb/server/handler"
	"github.com/crackeer/goweb/server/middleware"
	"github.com/gin-gonic/gin"
)

func Run() error {

	router := gin.New()
	config := container.GetConfig()
	gin.SetMode(gin.DebugMode)
	router.Any("page/login", middleware.InitPage(), middleware.RenderPage(), handler.RenderLogin)
	router.Any("page/logout", middleware.InitPage(), middleware.RenderPage(), handler.RenderLogout)

	router.GET("share/:code", middleware.InitPage(), middleware.RenderPage(), handler.RenderShare)
	router.Use(middleware.Login())
	router.StaticFile("database", config.Sqlite3DatabaseFile)
	api := router.Group("api")
	api.POST("object/update", handler.UpdateObject)
	api.POST("object/upload", handler.UploadObject)
	api.POST("object/append", handler.AppendObject)
	api.POST("object/delete", handler.DeleteObject)
	api.POST("object/share", handler.ShareObject)

	router.GET("image/:id", handler.RenderImage)

	page := router.Group("page", middleware.InitPage(), middleware.RenderPage())
	page.GET("index", handler.RenderIndex)
	page.GET("link/list", handler.RenderLinkList)
	page.GET("link/edit", handler.RenderEditLink)
	page.GET("diary/list", handler.RenderDiaryList)
	page.GET("diary/detail", handler.RenderMarkdown)
	page.GET("diary/edit", handler.RenderEditMarkdown)
	page.GET("markdown/detail", handler.RenderMarkdown)
	page.GET("markdown/create", handler.RenderCreateMarkdown)
	page.GET("markdown/list", handler.RenderMarkdownList)
	page.GET("markdown/edit", handler.RenderEditMarkdown)

	page.GET("code/create", handler.RenderCreateCode)
	page.GET("code/list", handler.RenderCodeList)
	page.GET("code/edit", handler.RenderEditCode)
	page.GET("code/detail", handler.RenderCode)
	router.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/page/index")
	})

	return router.Run(fmt.Sprintf(":%d", config.Port))
}
