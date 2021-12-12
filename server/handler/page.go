package handler

import (
	"net/http"

	"github.com/crackeer/goweb/container"
	objectService "github.com/crackeer/goweb/service/object"
	pageService "github.com/crackeer/goweb/service/page"
	"github.com/gin-gonic/gin"
)

// GetLinkList
//  @param ctx
func RenderLinkPage(ctx *gin.Context) {
	list := objectService.GetAllLinkList()

	page := pageService.NewPage("书签列表", list, container.GetSkeletionTemplatePath())

	html := page.Render(container.GetFullTemplatePath("links"))

	ctx.Data(http.StatusOK, "text/html", []byte(html))
}
