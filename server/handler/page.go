package handler

import (
	"net/http"

	"github.com/crackeer/goweb/common"
	"github.com/crackeer/goweb/container"
	"github.com/crackeer/goweb/model"
	objectService "github.com/crackeer/goweb/service/object"
	pageService "github.com/crackeer/goweb/service/page"
	"github.com/gin-gonic/gin"
)

// RenderIndex
//  @param ctx
func RenderIndex(ctx *gin.Context) {

	date := common.GetNowDate()
	obj, _ := model.GetTheOne(model.TypeMD, model.TagDiary, date)
	page := pageService.NewPage(ctx, "首页", map[string]interface{}{
		"date":  date,
		"diary": obj.ToMap(),
	})

	html := page.Render(container.GetFullTemplatePath("index"))

	ctx.Data(http.StatusOK, "text/html", []byte(html))
}

// RenderLinkList
//  @param ctx
func RenderLinkList(ctx *gin.Context) {
	list := objectService.GetAllLinkList()

	page := pageService.NewPage(ctx, "书签列表", list)

	html := page.Render(container.GetFullTemplatePath("links"))

	ctx.Data(http.StatusOK, "text/html", []byte(html))
}

// RenderLinkList
//  @param ctx
func RenderEditLink(ctx *gin.Context) {
	list := objectService.GetAllLinkList()

	page := pageService.NewPage(ctx, "修改书签", list)

	html := page.Render(container.GetFullTemplatePath("edit_links"))

	ctx.Data(http.StatusOK, "text/html", []byte(html))
}
