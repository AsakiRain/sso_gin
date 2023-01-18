package middleware

import (
	"net/http"
	"sso_gin/utils"

	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			ctx.Abort()
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    40101,
				"message": "token不存在",
			})
			return
		}
		// 校验token，只要出错直接拒绝请求
		_, err := utils.ParseToken(auth)
		if err != nil {
			ctx.Abort()
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    40102,
				"message": err.Error(),
			})
			return
		}
		ctx.Next()
	}
}
