package database

import (
	"github.com/crackeer/gopkg/ginhelper"
	"github.com/crackeer/goweb/container"
	"github.com/crackeer/goweb/model"
	"github.com/gin-gonic/gin"
)

// Exec ...
//  @param ctx
func Exec(ctx *gin.Context) {
	bytes, err := ctx.GetRawData()
	if err != nil {
		ginhelper.Failure(ctx, ginhelper.CodeDefaultError, err.Error())
		return
	}
	err = model.ExecSQL(container.GetDatabase(), string(bytes))
	if err != nil {
		ginhelper.Failure(ctx, ginhelper.CodeDefaultError, err.Error())
		return
	}
	ginhelper.Success(ctx, map[string]interface{}{
		"sql": string(bytes),
	})
}
