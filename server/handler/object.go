package handler

import (
	"net/http"

	"github.com/crackeer/goweb/model"
	"github.com/gin-gonic/gin"
)

// UpdateObject
//  @param ctx
func UpdateObject(ctx *gin.Context) {
	object := &model.Object{}
	if err := ctx.ShouldBindJSON(object); err != nil {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}

	if len(object.Title) < 1 {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code":    -2,
			"message": "title是必须的",
		})
		return
	}

	if err := object.Update(); err != nil {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"data": object.ToMap(),
	})
}

// DeleteObject
//  @param ctx
func DeleteObject(ctx *gin.Context) {
	object := &model.Object{}
	if err := ctx.ShouldBindJSON(object); err != nil {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}

	if object.ID < 1 {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code":    -2,
			"message": "id不能小于0",
		})
		return
	}

	model.DeleteObjectByID(object.ID)

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "删除成功",
		"data":    nil,
	})
}
