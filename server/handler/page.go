package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/crackeer/goweb/common"
	"github.com/crackeer/goweb/container"
	"github.com/crackeer/goweb/define"
	"github.com/crackeer/goweb/model"
	objectService "github.com/crackeer/goweb/service/object"
	pageService "github.com/crackeer/goweb/service/page"
	"github.com/gin-gonic/gin"
)

func getPager(ctx *gin.Context) *pageService.Page {
	val, exists := ctx.Get("renderer")
	if exists {
		if pager, ok := val.(*pageService.Page); ok {
			return pager
		}
	}
	return nil
}

func setData(ctx *gin.Context, data interface{}) {
	pager := getPager(ctx)
	if pager != nil {
		pager.SetData(data)
	}
}

func setTitle(ctx *gin.Context, title string) {
	pager := getPager(ctx)
	if pager != nil {
		pager.SetTitle(title)
	}
}

// RenderIndex
//  @param ctx
func RenderLogin(ctx *gin.Context) {
	conf := container.GetConfig()
	err := ""
	if ctx.Request.Method == http.MethodPost {
		password := common.MD5(ctx.PostForm("password"))
		if conf.PasswordMD5 == password {
			ctx.SetCookie(define.TokenKey, common.MD5(conf.PasswordMD5), 3600*24*30, "/", "", true, true)
			return
		}
		err = "密码错误"
	}

	date := common.GetNowDate()
	setData(ctx, map[string]interface{}{
		"date":  date,
		"error": err,
	})
}

// RenderIndex
//  @param ctx
func RenderIndex(ctx *gin.Context) {

	date := common.GetNowDate()
	obj, _ := model.GetTheDiary(model.TypeDiary, date)
	setData(ctx, map[string]interface{}{
		"date":   date,
		"object": obj.ToMap(),
	})
}

// RenderLinkList
//  @param ctx
func RenderLinkList(ctx *gin.Context) {
	setData(ctx, objectService.GetAllLinkList())
}

// RenderLinkList
//  @param ctx
func RenderEditLink(ctx *gin.Context) {
	setData(ctx, objectService.GetAllLinkList())
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

	setData(ctx, list)
}

// RenderDiaryList
//  @param ctx
func RenderMarkdown(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "0")

	val, _ := strconv.Atoi(id)
	object, _ := model.GetObjectByID(int64(val))

	setData(ctx, map[string]interface{}{
		"object": object.ToMap(),
	})
}

// RenderCreateMarkdown
//  @param ctx
func RenderCreateMarkdown(ctx *gin.Context) {

	tags, _ := model.GetTags(model.TypeMD)

	setData(ctx, tags)
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
	setTitle(ctx, title)
	setData(ctx, map[string]interface{}{
		"object": object.ToMap(),
		"tags":   tags,
	})
}

// RenderMarkdownList
//  @param ctx
func RenderMarkdownList(ctx *gin.Context) {

	tags, _ := model.GetTags(model.TypeMD)
	if len(tags) < 1 {
		setData(ctx, map[string]interface{}{
			"tags":  tags,
			"list":  []map[string]interface{}{},
			"tag":   "",
			"page":  1,
			"total": 0,
		})
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

	setData(ctx, map[string]interface{}{
		"tags":  tags,
		"list":  list,
		"page":  val,
		"tag":   tag,
		"total": total,
	})

}

// RenderCreateMarkdown
//  @param ctx
func RenderCreateCode(ctx *gin.Context) {

	lang := ctx.DefaultQuery("tag", "go")

	config := container.GetConfig()
	setData(ctx, map[string]interface{}{
		"tags": config.CodeLanguages,
		"tag":  lang,
	})
}

// RenderCreateMarkdown
//  @param ctx
func RenderEditCode(ctx *gin.Context) {

	id := ctx.DefaultQuery("id", "0")

	val, _ := strconv.Atoi(id)
	object, _ := model.GetObjectByID(int64(val))
	config := container.GetConfig()
	title := fmt.Sprintf("修改代码 - %s", object.Title)
	setTitle(ctx, title)
	setData(ctx, map[string]interface{}{
		"object": object.ToMap(),
		"tags":   config.CodeLanguages,
	})

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
	setData(ctx, map[string]interface{}{
		"tags":  config.CodeLanguages,
		"list":  list,
		"page":  val,
		"tag":   tag,
		"total": total,
	})

}

// RenderCode
//  @param ctx
func RenderCode(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "0")

	val, _ := strconv.Atoi(id)
	object, _ := model.GetObjectByID(int64(val))

	setData(ctx, map[string]interface{}{
		"object": object.ToMap(),
	})
}
