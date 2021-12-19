package server

import (
	"fmt"
	"net/http"

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
	api.POST("object/delete", handler.DeleteObject)

	page := router.Group("page")
	page.GET("index", handler.RenderIndex)
	page.GET("link/list", handler.RenderLinkList)
	page.GET("link/edit", handler.RenderEditLink)
	page.GET("diary/list", handler.RenderDiaryList)
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
