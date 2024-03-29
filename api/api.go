package api

import (
	api_email "sso_gin/api/email"
	api_login "sso_gin/api/login"
	api_reg "sso_gin/api/reg"
	api_user "sso_gin/api/user"
	"sso_gin/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	reg := router.Group("/reg/flow", middleware.FlowCheck())
	{
		reg.GET("/0", api_reg.HandleStepStart)
		reg.POST("/1", api_reg.HandleStepTOS)
		reg.POST("/2", api_reg.HandleStepEmail)
		reg.POST("/3", api_reg.HandleStepAccount)
		reg.POST("/4/start", api_reg.HandleMsStart)
		reg.POST("/4/link", api_reg.HandleGenerateLink)
		reg.POST("/4/status", api_reg.HandleMsStatus)
		reg.POST("/4", api_reg.HandleStepMs)
		reg.POST("/5/code", api_reg.HandleGetQqCode)
		reg.POST("/5/check", api_reg.HandleCheckQqCode)
		reg.POST("/5", api_reg.HandleStepQq)
		reg.GET("/6", api_reg.HandleStepPreference)
		reg.GET("/7", api_reg.HandleStepDone)
	}
	user := router.Group("/user", middleware.JwtAuth())
	{
		user.GET("/info", api_user.HandleUserInfo)
	}
	email := router.Group("/email")
	{
		email.POST("/code", api_email.HandleSendCode)
	}
	router.POST("/login", api_login.HandleLogin)
}
