package handler

import (
	"strconv"

	"github.com/crackeer/goweb/model"
	"github.com/gin-gonic/gin"
)

// RenderIndex
//  @param ctx
func RenderShare(ctx *gin.Context) {

	shareCode := ctx.Param("code")
	val, _ := strconv.Atoi(shareCode)
	object, _ := model.GetObjectByID(int64(val))
	setData(ctx, map[string]interface{}{
		"object": object.ToMap(),
	})
	setTitle(ctx, "分享："+object.Title)
	setPageType(ctx, "share")
	setTPLFile(ctx, "share")
}
