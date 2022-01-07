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

func setPageType(ctx *gin.Context, pageType string) {
	pager := getPager(ctx)
	if pager != nil {
		pager.SetPageType(pageType)
	}
}

func setTitle(ctx *gin.Context, title string) {
	pager := getPager(ctx)
	if pager != nil {
		pager.SetTitle(title)
	}
}

func setTPLFile(ctx *gin.Context, title string) {
	pager := getPager(ctx)
	if pager != nil {
		pager.SetTPLFile(container.GetFullTemplatePath(title))
	}
}

// RenderLogout
//  @param ctx
func RenderLogout(ctx *gin.Context) {
	conf := container.GetConfig()
	ctx.SetCookie(define.TokenKey, "", -1, "/", conf.Domain, false, false)
}

// RenderIndex
//  @param ctx
func RenderLogin(ctx *gin.Context) {
	conf := container.GetConfig()
	err := ""
	if ctx.Request.Method == http.MethodPost {
		password := common.MD5(ctx.PostForm("password"))
		if conf.PasswordMD5 == password {
			ctx.SetCookie(define.TokenKey, common.MD5(conf.PasswordMD5), 3600*24*30, "/", conf.Domain, false, false)
			setTPLFile(ctx, "login/success")
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
