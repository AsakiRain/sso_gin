package api_reg

import (
	"fmt"
	"log"
	"net/http"
	"sso_gin/db"
	"sso_gin/model"
	"sso_gin/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	uuid "github.com/satori/go.uuid"
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
		})
		return
	}
	email := stepEmailForm.Email
	code := strings.ToUpper(stepEmailForm.Code)
	serial := stepEmailForm.Serial

	valid := utils.CheckCode(email, code)
	if !valid {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "验证码错误",
		})
		return
	}
	cacheKey := fmt.Sprintf("email_captcha_%s", code)
	CACHE.Delete(cacheKey)

	uuidV4, err := uuid.NewV4()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "服务器错误",
		})
		log.Printf("未能产生uuid：%v", err)
		return
	}
	state := uuidV4.String()
	var regFlow model.RegFlow
	updateForm := map[string]interface{}{
		"step":     2,
		"ms_state": state,
	}
	MYSQL.Model(&regFlow).Where("serial = ?", serial).Updates(updateForm)

	ctx.JSON(http.StatusOK, gin.H{
		"code":        200,
		"message":     "验证成功",
		"url":         "/reg/flow/3",
		"link_start":  utils.GenerateLinkStart(state),
		"link_remake": utils.GenerateLinkRemake(),
	})
}
