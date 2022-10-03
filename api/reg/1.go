package api_reg

import (
	"net/http"
	"sso_gin/db"
	"sso_gin/model"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func HandleStepTOS(ctx *gin.Context) {
	MYSQL := *db.MYSQL
	var stepTOSForm model.StepTOSForm
	err := ctx.ShouldBindBodyWith(&stepTOSForm, binding.JSON)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
			"detail":  err.Error(),
		})
		return
	}
	acceptTos := stepTOSForm.AcceptTos
	serial := stepTOSForm.Serial

	if !*acceptTos {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "请同意服务条款",
			"detail":  "我叼你妈，你直接发包是吧",
		})
		return
	}
	var regFlow model.RegFlow
	MYSQL.Model(&regFlow).Where("serial = ?", serial).Update("step", 1)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "验证成功",
		"url":     "/reg/flow/2",
	})
}
