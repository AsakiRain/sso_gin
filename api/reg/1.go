package api_reg

import (
	"fmt"
	"net/http"
	"sso_gin/db"
	"sso_gin/model"

	"github.com/gin-gonic/gin"
)

func HandleStepTOS(ctx *gin.Context) {
	MYSQL := *db.MYSQL
	var stepTOSForm model.StepTOSForm
	err := ctx.ShouldBindJSON(&stepTOSForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}
	acceptTos := stepTOSForm.AcceptTos
	serial := stepTOSForm.Serial

	var regFlow model.RegFlow
	result := MYSQL.First(&regFlow, "serial = ?", serial)
	if result.Error != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "流程不存在",
			"url":     "/reg/flow/0",
		})
		return
	}
	if regFlow.Step != 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "流程错误",
			"url":     fmt.Sprintf("/reg/flow/%d", regFlow.Step+1),
		})
		return
	}
	if !*acceptTos {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "请同意服务条款",
			"detail":  "我叼你妈，你直接发包是吧",
		})
		return
	}
	MYSQL.Model(&regFlow).Where("serial = ?", serial).Update("step", 1)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "验证成功",
		"url":     "/reg/flow/2",
	})
}
