package test

import (
	"net/http"

	"github.com/crackeer/goweb/container"
	"github.com/gin-gonic/gin"
)

// RequestAPI
//  @param ctx
func RequestAPI(ctx *gin.Context) {
	response, err := container.APIRequestClient.Request("goweb/api_v1_list_object", map[string]interface{}{}, map[string]string{}, "")
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"response": string(response.OriginBody),
		"error":    err,
	})
}
