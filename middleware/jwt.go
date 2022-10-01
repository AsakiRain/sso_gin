package middleware

import (
	"net/http"
	"sso_gin/utils"

	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			c.Abort()
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    422,
				"message": "token不存在",
			})
			return
		}
		// 校验token，只要出错直接拒绝请求
		_, err := utils.ParseToken(auth)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    422,
				"message": err.Error(),
			})
			return
		}
		c.Next()
	}
}
