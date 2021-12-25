package handler

import (
	"fmt"
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
		"date":   date,
		"object": obj.ToMap(),
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

	page := pageService.NewPage(ctx, "书签管理", list)

	html := page.Render(container.GetFullTemplatePath("edit_link"))

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
	id := ctx.DefaultQuery("id", "0")

	val, _ := strconv.Atoi(id)
	object, _ := model.GetObjectByID(int64(val))

	pager := pageService.NewPage(ctx, object.Title, map[string]interface{}{
		"object": object.ToMap(),
	})

	html := pager.Render(container.GetFullTemplatePath("markdown"))

	ctx.Data(http.StatusOK, "text/html", []byte(html))
}

// RenderCreateMarkdown
//  @param ctx
func RenderCreateMarkdown(ctx *gin.Context) {

	tags, _ := model.GetTags(model.TypeMD)

	pager := pageService.NewPage(ctx, "创建markdown文档", tags)

	html := pager.Render(container.GetFullTemplatePath("create_markdown"))

	ctx.Data(http.StatusOK, "text/html", []byte(html))
}

// RenderEditMarkdown
//  @param ctx
func RenderEditMarkdown(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "0")

	val, _ := strconv.Atoi(id)
	object, _ := model.GetObjectByID(int64(val))
	tags, _ := model.GetTags(model.TypeMD)
	title := fmt.Sprintf("修改文档 - %s", object.Title)
	if object.Type == model.TypeDiary {
		title = fmt.Sprintf("修改日记 - %s", object.Title)
	}
	pager := pageService.NewPage(ctx, title, map[string]interface{}{
		"object": object.ToMap(),
		"tags":   tags,
	})

	html := pager.Render(container.GetFullTemplatePath("edit_markdown"))

	ctx.Data(http.StatusOK, "text/html", []byte(html))
}

// RenderMarkdownList
//  @param ctx
func RenderMarkdownList(ctx *gin.Context) {

	tags, _ := model.GetTags(model.TypeMD)
	if len(tags) < 1 {
		pager := pageService.NewPage(ctx, "markdown文档列表", map[string]interface{}{
			"tags":  tags,
			"list":  []map[string]interface{}{},
			"tag":   "",
			"page":  1,
			"total": 0,
		})
		html := pager.Render(container.GetFullTemplatePath("markdown_list"))
		ctx.Data(http.StatusOK, "text/html", []byte(html))
		return
	}
	page := ctx.DefaultQuery("page", "1")
	val, _ := strconv.Atoi(page)

	tag := ctx.DefaultQuery("tag", tags[0])
	objects, total, _ := model.GetObjectList(model.TypeMD, tag, int64(val))

	list := []map[string]interface{}{}
	for _, v := range objects {
		list = append(list, v.ToMap())
	}

	pager := pageService.NewPage(ctx, "markdown文档列表", map[string]interface{}{
		"tags":  tags,
		"list":  list,
		"page":  val,
		"tag":   tag,
		"total": total,
	})

	html := pager.Render(container.GetFullTemplatePath("markdown_list"))
	ctx.Data(http.StatusOK, "text/html", []byte(html))
}

// RenderCreateMarkdown
//  @param ctx
func RenderCreateCode(ctx *gin.Context) {

	lang := ctx.DefaultQuery("tag", "go")

	config := container.GetConfig()
	pager := pageService.NewPage(ctx, "创建代码片段", map[string]interface{}{
		"tags": config.CodeLanguages,
		"tag":  lang,
	})

	html := pager.Render(container.GetFullTemplatePath("create_code"))

	ctx.Data(http.StatusOK, "text/html", []byte(html))
}

// RenderCreateMarkdown
//  @param ctx
func RenderEditCode(ctx *gin.Context) {

	id := ctx.DefaultQuery("id", "0")

	val, _ := strconv.Atoi(id)
	object, _ := model.GetObjectByID(int64(val))
	config := container.GetConfig()
	title := fmt.Sprintf("修改代码 - %s", object.Title)
	pager := pageService.NewPage(ctx, title, map[string]interface{}{
		"object": object.ToMap(),
		"tags":   config.CodeLanguages,
	})

	html := pager.Render(container.GetFullTemplatePath("edit_code"))

	ctx.Data(http.StatusOK, "text/html", []byte(html))
}

// RenderCreateMarkdown
//  @param ctx
func RenderCodeList(ctx *gin.Context) {

	page := ctx.DefaultQuery("page", "1")
	val, _ := strconv.Atoi(page)

	tag := ctx.DefaultQuery("tag", "go")
	objects, total, _ := model.GetObjectList(model.TypeCode, tag, int64(val))
	list := []map[string]interface{}{}
	for _, v := range objects {
		list = append(list, v.ToMap())
	}
	config := container.GetConfig()
	pager := pageService.NewPage(ctx, "代码列表", map[string]interface{}{
		"tags":  config.CodeLanguages,
		"list":  list,
		"page":  val,
		"tag":   tag,
		"total": total,
	})

	html := pager.Render(container.GetFullTemplatePath("code_list"))
	ctx.Data(http.StatusOK, "text/html", []byte(html))
}

// RenderDiaryList
//  @param ctx
func RenderCode(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "0")

	val, _ := strconv.Atoi(id)
	object, _ := model.GetObjectByID(int64(val))

	pager := pageService.NewPage(ctx, object.Title, map[string]interface{}{
		"object": object.ToMap(),
	})

	html := pager.Render(container.GetFullTemplatePath("code"))

	ctx.Data(http.StatusOK, "text/html", []byte(html))
}
