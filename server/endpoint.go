package server

import (
	"fmt"

	"github.com/crackeer/goweb/container"
	"github.com/crackeer/goweb/server/handler"
	"github.com/gin-gonic/gin"
)

func Run() error {

	router := gin.New()
	config := container.GetConfig()
	gin.SetMode(gin.DebugMode)

	api := router.Group("api")
	api.POST("object/update", handler.UpdateObject)

	page := router.Group("page")
	page.GET("index", handler.RenderIndex)
	page.GET("link/list", handler.RenderLinkList)
	page.GET("link/edit", handler.RenderEditLink)
	page.GET("diary/list", handler.RenderDiaryList)
	page.GET("markdown/detail/:id", handler.RenderMarkdown)
	page.GET("markdown/create", handler.RenderCreateMarkdown)
	page.GET("markdown/list", handler.RenderMarkdownList)
	router.NoRoute(handler.RenderIndex)

	return router.Run(fmt.Sprintf(":%d", config.Port))
}
