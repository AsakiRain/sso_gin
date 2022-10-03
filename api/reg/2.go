package api_reg

import (
	"fmt"
	"net/http"
	"sso_gin/db"
	"sso_gin/model"
	"sso_gin/utils"
	"strings"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func HandleStepEmail(ctx *gin.Context) {
	CACHE := *db.CACHE
	var stepEmailForm model.StepEmailForm
	err := ctx.ShouldBindJSON(&stepEmailForm)
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

	//应当检测流程是否存在
	fmt.Printf("Serial(暂时不用):%s", serial)
	//

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

	//应当更新数据库
	//
	state := uuid.NewV4().String()

	ctx.JSON(http.StatusOK, gin.H{
		"code":       200,
		"message":    "验证成功",
		"url":        "/reg/flow/3",
		"link_start": utils.GenerateLinkStart(state),
	})
}
