package table

import (
	"fmt"

	"github.com/crackeer/gopkg/ginhelper"
	"github.com/crackeer/gopkg/util"
	"github.com/crackeer/goweb/container"
	"github.com/crackeer/goweb/model"
	"github.com/gin-gonic/gin"
)

// Update ...
//  @param ctx
func Update(ctx *gin.Context) {
	tableName := ctx.Param("table")
	tableObj, err := model.NewTable(container.GetDatabase(), tableName)
	if err != nil {
		ginhelper.Failure(ctx, -1, err.Error())
		return
	}
	params := ginhelper.AllParams(ctx)

	primaryKey := tableObj.GetPrimaryKey()
	primaryKeyValue := util.LoadMap(params).GetString(primaryKey, "")
	if len(primaryKeyValue) < 1 {
		ginhelper.Failure(ctx, -1, fmt.Sprintf("primaryKey `%s` value not found", primaryKey))
		return
	}

	result := tableObj.Update(map[string]interface{}{
		primaryKey: primaryKeyValue,
	}, params)
	if result < 1 {
		ginhelper.Failure(ctx, -1, "not update")
		return
	}
	ginhelper.Success(ctx, params)
}
