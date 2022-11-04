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
			"data":    nil,
		})
		return
	}
	acceptTos := stepTOSForm.AcceptTos
	serial := stepTOSForm.Serial

	if !*acceptTos {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "你怎么直接发包啊（恼）",
			"data":    nil,
		})
		return
	}

	postForm := map[string]interface{}{
		"step":       1,
		"accept_tos": acceptTos,
	}
	MYSQL.Model(&model.RegFlow{}).Where("serial = ?", serial).Updates(postForm)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "验证成功",
		"data": map[string]interface{}{
			"url": "/reg/flow/2",
		},
	})
}
