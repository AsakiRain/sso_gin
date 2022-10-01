package api_email

import (
	"fmt"
	"net/http"
	"sso_gin/db"
	"sso_gin/model"
	"sso_gin/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type EmailForm struct {
	Email string `json:"email"`
}

func HandleSendCode(ctx *gin.Context) {
	MYSQL := *db.MYSQL
	CACHE := *db.CACHE
	var emailForm EmailForm
	err := ctx.ShouldBindJSON(&emailForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}
	var userInfo model.UserInfo
	result := MYSQL.First(&userInfo, "email = ?", emailForm.Email)
	if result.Error == nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "邮箱已经注册",
		})
		return
	}

	code := utils.RandomCode(6)
	cacheKey := fmt.Sprintf("email_captcha_%s", code)
	emailCaptcha := model.EmailCaptcha{
		Email:     emailForm.Email,
		Code:      code,
		ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
	}
	CACHE.Set(cacheKey, emailCaptcha, time.Minute*10)
	mailBody, err := utils.ParseTemplate("captcha.html", map[string]interface{}{"code": code, "ip": ctx.ClientIP()})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "发件失败",
			"detail":  err.Error(),
		})
		return
	}
	err = utils.SendMail(emailForm.Email, "注册验证码", mailBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "发件失败",
			"detail":  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "发送成功",
	})
}
