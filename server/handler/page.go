package handler

import (
	"net/http"
	"strconv"

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
	obj, _ := model.GetTheDiary(model.TypeDiary, date)
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

// RenderDiaryList
//  @param ctx
func RenderDiaryList(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")

	val, _ := strconv.Atoi(page)
	objects, _ := model.GetDiaryList(int64(val))

	list := []map[string]interface{}{}

	for _, v := range objects {
		list = append(list, v.ToMap())
	}

	pager := pageService.NewPage(ctx, "日记列表", list)

	html := pager.Render(container.GetFullTemplatePath("diary_list"))

	ctx.Data(http.StatusOK, "text/html", []byte(html))
}

// RenderDiaryList
//  @param ctx
func RenderMarkdown(ctx *gin.Context) {
	id := ctx.Param("id")

	val, _ := strconv.Atoi(id)
	object, _ := model.GetObjectByID(int64(val))

	pager := pageService.NewPage(ctx, object.Title, object.ToMap())

	html := pager.Render(container.GetFullTemplatePath("markdown"))

	ctx.Data(http.StatusOK, "text/html", []byte(html))
}

// RenderCreateMarkdown
//  @param ctx
func RenderCreateMarkdown(ctx *gin.Context) {

	tags, _ := model.GetTags(model.TypeMD)

	pager := pageService.NewPage(ctx, "创建文档", tags)

	html := pager.Render(container.GetFullTemplatePath("create_markdown"))

	ctx.Data(http.StatusOK, "text/html", []byte(html))
}

// RenderMarkdownList
//  @param ctx
func RenderMarkdownList(ctx *gin.Context) {

	tags, _ := model.GetTags(model.TypeMD)
	if len(tags) < 1 {
		pager := pageService.NewPage(ctx, "markdown文档列表", map[string]interface{}{
			"tags": tags,
			"list": []map[string]interface{}{},
			"tag":  "",
			"page": 1,
		})
		html := pager.Render(container.GetFullTemplatePath("markdown_list"))
		ctx.Data(http.StatusOK, "text/html", []byte(html))
		return
	}
	page := ctx.DefaultQuery("page", "1")
	val, _ := strconv.Atoi(page)

	tag := ctx.DefaultQuery("tag", tags[0])
	objects, _ := model.GetObjectList(model.TypeMD, tag, int64(val))

	list := []map[string]interface{}{}
	for _, v := range objects {
		list = append(list, v.ToMap())
	}

	pager := pageService.NewPage(ctx, "markdown文档列表", map[string]interface{}{
		"tags": tags,
		"list": list,
		"page": val,
		"tag":  tag,
	})

	html := pager.Render(container.GetFullTemplatePath("markdown_list"))
	ctx.Data(http.StatusOK, "text/html", []byte(html))
}
