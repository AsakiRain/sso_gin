package middleware

import (
	"fmt"
	"net/http"
	"sso_gin/db"
	"sso_gin/model"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func FlowCheck() gin.HandlerFunc {
	MYSQL := *db.MYSQL
	return func(ctx *gin.Context) {
		if ctx.Request.URL.Path == "/reg/flow/0" {
			ctx.Next()
			return
		}
		var serialForm model.SerialForm
		err := ctx.ShouldBindBodyWith(&serialForm, binding.JSON)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "没有提供流水号",
				"url":     "/reg/flow/0",
			})
			ctx.Abort()
			return
		}
		serial := serialForm.Serial

		var regFlow model.RegFlow
		result := MYSQL.First(&regFlow, "serial = ?", serial)
		if result.Error != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    422,
				"message": "流程不存在",
				"url":     "/reg/flow/0",
			})
			ctx.Abort()
			return
		}

		yourStep := ctx.Request.URL.Path[len("/reg/flow/"):]
		myStep := fmt.Sprintf("%d", regFlow.Step+1)
		if yourStep != myStep {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    422,
				"message": "流程错误",
				"url":     fmt.Sprintf("/reg/flow/%v", myStep),
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
