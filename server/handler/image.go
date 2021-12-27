package handler

import (
	"net/http"
	"strconv"

	"github.com/crackeer/goweb/model"
	"github.com/gin-gonic/gin"
)

// RenderImage
//  @param ctx
func RenderImage(ctx *gin.Context) {
	id := ctx.Param("id")
	val, _ := strconv.Atoi(id)
	object, _ := model.GetObjectByID(int64(val))

	ctx.Data(http.StatusOK, "image/jpeg", []byte(object.Content))
}
