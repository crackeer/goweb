package table

import (
	"github.com/crackeer/gopkg/ginhelper"
	"github.com/gin-gonic/gin"
)

// Create
//  @param ctx
func Create(ctx *gin.Context) {
	tableObj, err := getTableObject(ctx)
	if err != nil {
		ginhelper.Failure(ctx, -1, err.Error())
		return
	}
	params := ginhelper.AllParams(ctx)

	params, err = tableObj.Create(params)
	tableObj.GetPrimaryKey()
	if err != nil {
		ginhelper.Failure(ctx, -1, err.Error())
		return
	}
	ginhelper.Success(ctx, params)
}
