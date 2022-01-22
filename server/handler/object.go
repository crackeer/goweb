package handler

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/crackeer/goweb/common"
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

// UploadObject
//  @param ctx
func UploadObject(ctx *gin.Context) {
	header, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}

	tmpFile, err := header.Open()

	if err != nil {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}
	bytes, err := ioutil.ReadAll(tmpFile)
	if err != nil {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}

	object := &model.Object{
		Content: string(bytes),
		Type:    model.TypeImage,
		Title:   common.MD5(header.Filename),
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
		"data": map[string]interface{}{
			"id": object.ID,
		},
	})
}

type share struct {
	ID       int64 `json:"id"`
	Duration int64 `json:"duration"`
}

// ShareObject
//  @param ctx
func ShareObject(ctx *gin.Context) {

	object := &share{}
	if err := ctx.ShouldBindJSON(object); err != nil {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}

	var expire int64 = -1
	if object.Duration > 0 {
		expire = object.Duration + time.Now().Unix()
	}

	shareCode, err := model.ShareObject(object.ID, expire)
	if err != nil {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code":    -3,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "删除成功",
		"data": map[string]interface{}{
			"share_url": "/share/" + shareCode,
		},
	})
}
