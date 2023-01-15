package api_email

import (
	"fmt"
	"net/http"
	"sso_gin/db"
	"sso_gin/model"
	"sso_gin/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func HandleSendCode(ctx *gin.Context) {
	MYSQL := *db.MYSQL
	CACHE := *db.CACHE
	var emailForm model.EmailForm
	err := ctx.ShouldBindJSON(&emailForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    40001,
			"message": "参数错误",
			"data":    nil,
		})
		return
	}
	email := strings.ToLower(emailForm.Email)
	cdKey := fmt.Sprintf("email_cd_%s", email)
	x, found := CACHE.Get(cdKey)
	if found {
		cdAt := x.(int64)
		if cdAt > time.Now().Unix() {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    42207,
				"message": fmt.Sprintf("操作频繁，请在%d秒后重试", cdAt-time.Now().Unix()),
				"data":    nil,
			})
			return
		}
	}

	var userInfo model.UserInfo
	result := MYSQL.First(&userInfo, "email = ?", email)
	if result.Error == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    42205,
			"message": "邮箱已经注册",
			"data":    nil,
		})
		return
	}

	code := utils.RandomCode(6)
	cacheKey := fmt.Sprintf("email_captcha_%s", code)
	emailCaptcha := model.EmailCaptcha{
		Email:     email,
		Code:      code,
		ExpiresAt: time.Now().Add(time.Minute * 10).UnixMilli(),
	}
	CACHE.Set(cacheKey, emailCaptcha, time.Minute*10)
	CACHE.Set(cdKey, time.Now().Add(time.Second*60).Unix(), time.Second*60)
	startTime := time.Now()
	mailBody, err := utils.ParseTemplate("captcha.html", map[string]interface{}{"code": code, "ip": ctx.ClientIP()})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    50000,
			"message": fmt.Sprintf("发件失败：%s", err.Error()),
			"data":    nil,
		})
		return
	}
	err = utils.SendMail(email, "注册验证码", mailBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    50000,
			"message": fmt.Sprintf("发件失败：%s", err.Error()),
			"data":    nil,
		})
		return
	}
	endTime := time.Now()
	ctx.JSON(http.StatusOK, gin.H{
		"code":    20000,
		"message": fmt.Sprintf("发送成功，耗时%f秒", endTime.Sub(startTime).Seconds()),
		"data":    nil,
	})
}
