package table

import (
	"github.com/crackeer/gopkg/ginhelper"
	"github.com/crackeer/gopkg/util"
	"github.com/gin-gonic/gin"
)

// Distinct
//  @param ctx
func Distinct(ctx *gin.Context) {
	tableObj, err := getTableObject(ctx)
	if err != nil {
		ginhelper.Failure(ctx, -1, err.Error())
		return
	}
	params := ginhelper.AllParams(ctx)
	field := util.LoadMap(params).GetString("_field", "")
	if len(field) < 1 {
		ginhelper.Failure(ctx, -1, "_field_ required")
		return
	}
	delete(params, "_field")

	list := tableObj.Distinct(field, params)

	ginhelper.Success(ctx, list)
}
