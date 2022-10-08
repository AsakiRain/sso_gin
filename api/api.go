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
	reg := router.Group("/reg/flow", middleware.FlowCheck())
	{
		reg.GET("/0", api_reg.HandleStepStart)
		reg.POST("/1", api_reg.HandleStepTOS)
		reg.POST("/2", api_reg.HandleStepEmail)
		reg.POST("/3", api_reg.HandleStepAccount)
		reg.POST("/4", api_reg.HandleStepMs)
		reg.GET("/4", api_reg.HandleMsQuery)
		reg.POST("/5", api_reg.HandleStepQq)
		reg.GET("/6", api_reg.HandleStepPerference)
		reg.GET("/7", api_reg.HandleStepDone)
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
