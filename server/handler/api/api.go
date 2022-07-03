package api

import (
	"fmt"

	ginhelper "github.com/crackeer/gopkg/gin"
	"github.com/crackeer/gopkg/util"
	"github.com/crackeer/goweb/container"
	"github.com/gin-gonic/gin"
)

// RequestAPI ...
//  @param ctx
func RequestAPI(ctx *gin.Context) {
	service := ctx.Param("service")
	api := ctx.Param("api")
	params := ginhelper.AllParams(ctx)
	appConfig := container.GetAppConfig()
	response, err := container.APIRequestClient.Request(service+"/"+api, params, map[string]string{}, appConfig.Env)
	fmt.Println(response.Data)
	if err != nil {
		ginhelper.Failure(ctx, -1, err.Error())
	} else {
		var apiData interface{}
		if err := util.Unmarshal(response.Data, &apiData); err != nil {
			apiData = response.Data
		}
		ginhelper.Success(ctx, apiData)
	}
}
