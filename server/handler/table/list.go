package table

import (
	"fmt"
	"net/http"

	"github.com/crackeer/goweb/container"
	"github.com/crackeer/goweb/model"
	"github.com/gin-gonic/gin"
)

// GetList
//  @param ctx
func GetList(ctx *gin.Context) {
	tableName := ctx.Param("table")
	fmt.Println(tableName)
	tableObj, err := model.NewTable(container.GetDatabase(), tableName)
	if err != nil {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}
	list := tableObj.Query(map[string]interface{}{})
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": -1,
		"data": list,
	})
}
