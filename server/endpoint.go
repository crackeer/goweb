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
	page.GET("links", handler.RenderLinkList)
	page.GET("edit_links", handler.RenderEditLink)

	return router.Run(fmt.Sprintf(":%d", config.Port))
}
