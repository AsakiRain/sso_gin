package api_reg

import (
	"fmt"
	"net/http"
	"sso_gin/db"
	"sso_gin/model"
	"sso_gin/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func HandleStepEmail(ctx *gin.Context) {
	MYSQL := *db.MYSQL
	CACHE := *db.CACHE
	var stepEmailForm model.StepEmailForm
	err := ctx.ShouldBindBodyWith(&stepEmailForm, binding.JSON)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
			"data":    nil,
		})
		return
	}
	email := strings.ToLower(stepEmailForm.Email)
	code := strings.ToUpper(stepEmailForm.Code)
	serial := stepEmailForm.Serial

	valid := utils.CheckCode(email, code)
	if !valid {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    422,
			"message": "验证码错误",
			"data":    nil,
		})
		return
	}
	cacheKey := fmt.Sprintf("email_captcha_%s", code)
	CACHE.Delete(cacheKey)

	updateForm := map[string]interface{}{
		"step":  2,
		"email": email,
	}
	MYSQL.Model(&model.RegFlow{}).Where("serial = ?", serial).Updates(updateForm)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "验证成功",
		"data": map[string]interface{}{
			"url": "/reg/flow/3",
		},
	})
}
