package table

import (
	"fmt"

	"github.com/crackeer/goweb/container"
	"github.com/crackeer/goweb/model"
	"github.com/gin-gonic/gin"
)

func getTableObject(ctx *gin.Context) (*model.Table, error) {
	tableName := ctx.Param("table")
	database := ctx.Param("database")
	fmt.Println(tableName, database)

	return model.NewTable(container.GetWriteDB(database), tableName)
}
