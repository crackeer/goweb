package middleware

import (
	"net/http"

	"github.com/crackeer/goweb/common"
	"github.com/crackeer/goweb/container"
	"github.com/crackeer/goweb/define"
	"github.com/gin-gonic/gin"
)

// Login
//  @return gin.HandlerFunc
func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie(define.TokenKey)

		if err != nil {
			ctx.Redirect(http.StatusTemporaryRedirect, getLoginURL(ctx))
			ctx.Abort()
		}
		conf := container.GetConfig()
		if common.MD5(conf.PasswordMD5) != token {
			ctx.Redirect(http.StatusTemporaryRedirect, getLoginURL(ctx))
			ctx.Abort()
		}
	}
}

func getLoginURL(ctx *gin.Context) string {
	return "/page/login"
}
