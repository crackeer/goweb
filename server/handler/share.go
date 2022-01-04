package handler

import (
	"strconv"
	"time"

	"github.com/crackeer/goweb/model"
	"github.com/gin-gonic/gin"
)

// RenderIndex
//  @param ctx
func RenderShare(ctx *gin.Context) {

	shareCode := ctx.Param("code")

	shareObject, err := model.GetObjectByTag(model.TypeShare, shareCode)

	setPageType(ctx, "share")
	setTPLFile(ctx, "share")
	setTitle(ctx, "暂无法打开分享链接")
	if err != nil || shareObject == nil || shareObject.ID < 1 {
		setData(ctx, map[string]interface{}{
			"error":   true,
			"message": "该分享不存在，或已经被移除",
		})
		return
	}

	expire, _ := strconv.Atoi(shareObject.Content)

	if expire > 0 && time.Now().Unix() > int64(expire) {
		setData(ctx, map[string]interface{}{
			"error":   true,
			"message": "对不起，分享的内容已经过期",
		})
		return
	}

	shareID, _ := strconv.Atoi(shareObject.Title)
	object, _ := model.GetObjectByID(int64(shareID))
	if object == nil || object.ID < 1 {
		setData(ctx, map[string]interface{}{
			"error":   true,
			"message": "来晚啦，分享的内容已经被删除了",
		})
		return
	}

	setData(ctx, map[string]interface{}{
		"error":  false,
		"object": object.ToMap(),
	})

	setTitle(ctx, "分享："+object.Title)
}
