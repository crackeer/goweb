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
	router.Any("/login", middleware.InitPage(), middleware.RenderPage(), handler.RenderLogin)
	router.Any("/logout", middleware.InitPage(), middleware.RenderPage(), handler.RenderLogout)

	router.GET("share/:code", middleware.InitPage(), middleware.RenderPage(), handler.RenderShare)
	router.Use(middleware.Login())
	router.StaticFile("database", config.Database)
	api := router.Group("api")
	
	api.POST("object/update", handler.UpdateObject)
	api.POST("object/upload", handler.UploadObject)
	api.POST("object/delete", handler.DeleteObject)
	api.POST("object/share", handler.ShareObject)

	router.GET("image/:id", handler.RenderImage)

	page := router.Group("", middleware.InitPage(), middleware.RenderPage())
	page.GET("link/list", handler.RenderLinkList)
	page.GET("link/edit", handler.RenderEditLink)
	page.GET("markdown/detail", handler.RenderMarkdown)
	page.GET("markdown/create", handler.RenderCreateMarkdown)
	page.GET("markdown/list", handler.RenderMarkdownList)
	page.GET("markdown/edit", handler.RenderEditMarkdown)

	router.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/markdown/list")
	})

	return router.Run(fmt.Sprintf(":%d", config.Port))
}
