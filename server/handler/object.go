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
