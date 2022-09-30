package api

import "github.com/gin-gonic/gin"

func SetupRouter(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.POST("/register", HandleRegister)
	}
}
