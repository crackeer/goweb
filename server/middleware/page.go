package middleware

import (
	"net/http"
	"strings"

	"github.com/crackeer/goweb/container"
	"github.com/crackeer/goweb/define"
	pageService "github.com/crackeer/goweb/service/page"
	"github.com/gin-gonic/gin"
)

// InitPage
//  @return gin.HandlerFunc
func InitPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pager := pageService.DefaultPage(ctx)
		conf := getPageConfig(ctx)
		pager.SetPageType(conf.Type)
		pager.SetTPLFile(container.GetFullTemplatePath(conf.TPL))
		pager.SetTitle(conf.Title)
		ctx.Set("renderer", pager)
	}
}

// RenderPage
//  @return gin.HandlerFunc
func RenderPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		val, exists := ctx.Get("renderer")
		if exists {
			if pager, ok := val.(*pageService.Page); ok {
				ctx.Data(http.StatusOK, "text/html", []byte(pager.Render()))
			}
		}
	}
}

func getPageConfig(ctx *gin.Context) define.PageConf {
	parts := strings.Split(strings.Trim(ctx.Request.URL.Path, "/"), "/")
	config := container.GetConfig()
	page := "index"
	if len(parts) >= 2 {
		page = strings.Join(parts[1:], "/")
	}

	conf, exists := config.Page[page]

	if !exists {
		return define.PageConf{
			Type: page,
			TPL:  page,
		}
	}

	if len(conf.TPL) < 1 {
		conf.TPL = page
	}

	if len(conf.Type) < 1 {
		conf.Type = parts[1]
	}

	return conf
}
