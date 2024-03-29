package middleware

import (
	"fmt"
	"log"
	"net/http"
	"sso_gin/db"
	"sso_gin/model"
	"sso_gin/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var NO_CHECK_SERIAL_LIST = []string{
	"/reg/flow/0",
	"/reg/flow/5/check",
}

var NO_CHECK_STEP_LIST = []string{
	"/reg/flow/4/query",
}

func FlowCheck() gin.HandlerFunc {
	MYSQL := *db.MYSQL
	return func(ctx *gin.Context) {
		if utils.Contains(NO_CHECK_SERIAL_LIST, ctx.FullPath()) {
			ctx.Next()
			return
		}
		var serialForm model.SerialForm
		err := ctx.ShouldBindBodyWith(&serialForm, binding.JSON)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":    40001,
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
			ctx.JSON(http.StatusOK, gin.H{
				"code":    42208,
				"message": "流程不存在",
				"data": map[string]interface{}{
					"url": "/reg",
				},
			})
			ctx.Abort()
			return
		}
		/////////////////////////////////
		// SHOULD REMOVE IN PRODUCTION //
		/////////////////////////////////
		if ctx.Request.URL.Query().Get("dev") == "true" {
			log.Println("Developer passed step check.")
			ctx.Next()
			return
		}
		yourStep := ctx.Request.URL.Path[len("/reg/flow/") : len("/reg/flow/")+1]
		if regFlow.Step >= 3 && utils.Contains(NO_CHECK_STEP_LIST, ctx.FullPath()) {
			ctx.Next()
			return
		}

		myStep := fmt.Sprintf("%d", regFlow.Step+1)
		if yourStep != myStep {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    42210,
				"message": fmt.Sprintf("流程错误，你应该在第 %v 步", myStep),
				"data": map[string]interface{}{
					"url": fmt.Sprintf("/reg/flow/%v", myStep),
				},
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
