package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Login
//  @param ctx
func Login(ctx *gin.Context) {

	fmt.Println(ctx.PostForm("password"))
}
