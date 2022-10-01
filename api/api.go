package api

import (
	api_email "sso_gin/api/email"
	api_reg "sso_gin/api/reg"
	api_user "sso_gin/api/user"
	"sso_gin/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	router.POST("/register", HandleRegister)
	reg := router.Group("/reg/flow")
	{
		reg.GET("/0", api_reg.HandleStepStart)
		reg.POST("/2", api_reg.HandleStepEmail)
	}
	user := router.Group("/user", middleware.JwtAuth())
	{
		user.GET("/info", api_user.HandleUserInfo)
	}
	email := router.Group("/email")
	{
		email.GET("/code", api_email.HandleSendCode)
	}
}