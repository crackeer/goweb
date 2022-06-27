package database

import (
	"github.com/crackeer/gopkg/ginhelper"
	"github.com/crackeer/goweb/container"
	"github.com/crackeer/goweb/model"
	"github.com/gin-gonic/gin"
)

// Tables ...
//  @param ctx
func Tables(ctx *gin.Context) {
	list := model.AllTables(container.GetDatabase())
	ginhelper.Success(ctx, list)
}
