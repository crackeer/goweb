package table

import (
	"github.com/crackeer/goweb/container"
	"github.com/crackeer/goweb/model"
	"github.com/gin-gonic/gin"
)

func getTableObject(ctx *gin.Context) (*model.Table, error) {
	tableName := ctx.Param("table")
	database := ctx.Param("database")

	return model.NewTable(container.GetDatabase(database), tableName)
}
