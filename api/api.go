package api

import (
	api_user "sso_gin/api/user"
	"sso_gin/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	router.POST("/register", HandleRegister)
	user := router.Group("/user", middleware.JwtAuth())
	{
		user.GET("/info", api_user.HandleUserInfo)
	}
}
