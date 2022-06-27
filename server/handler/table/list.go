package table

import (
	"github.com/crackeer/gopkg/ginhelper"
	"github.com/crackeer/gopkg/util"
	"github.com/gin-gonic/gin"
)

// List
//  @param ctx
func List(ctx *gin.Context) {
	tableObj, err := getTableObject(ctx)
	if err != nil {
		ginhelper.Failure(ctx, -1, err.Error())
		return
	}
	params := ginhelper.AllParams(ctx)
	page := util.LoadMap(params).GetInt64("_page_", 1)
	pageSize := util.LoadMap(params).GetInt64("_page_size_", 20)
	delete(params, "_page_")
	delete(params, "_page_size_")

	list := tableObj.GetPageList(params, page, pageSize)
	total := tableObj.Count(params)

	totalPage := total / pageSize
	if total%pageSize != 0 {
		totalPage = 1 + totalPage
	}

	ginhelper.Success(ctx, map[string]interface{}{
		"list":       list,
		"page_size":  pageSize,
		"page":       page,
		"total":      total,
		"total_page": totalPage,
	})
}
