package table

import (
	"github.com/crackeer/gopkg/ginhelper"
	"github.com/crackeer/goweb/container"
	"github.com/crackeer/goweb/model"
	"github.com/gin-gonic/gin"
)

// GetList
//  @param ctx
func Query(ctx *gin.Context) {
	tableName := ctx.Param("table")
	params := ginhelper.AllParams(ctx)

	tableObj, err := model.NewTable(container.GetDatabase(), tableName)
	if err != nil {
		ginhelper.Failure(ctx, -1, err.Error())
		return
	}
	list := tableObj.Query(params, 100)
	ginhelper.Success(ctx, list)
}
