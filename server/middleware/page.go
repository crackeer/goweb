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
	page := strings.Trim(ctx.Request.URL.Path, "/")

	conf := define.PageConf{
		Type: page,
		TPL:  page,
	}

	parts := strings.Split(page, "/")
	if len(parts) > 0 {
		conf.Type = parts[0]
	}

	return conf
}
