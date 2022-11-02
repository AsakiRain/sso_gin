package api_reg

import (
	"log"
	"net/http"
	"sso_gin/constant"
	"sso_gin/db"
	"sso_gin/model"
	"sso_gin/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func HandleStepAccount(ctx *gin.Context) {
	MYSQL := *db.MYSQL
	var stepAccountForm model.StepAccountForm
	err := ctx.ShouldBindBodyWith(&stepAccountForm, binding.JSON)
	if err != nil {
		ctx.JSON(400, gin.H{
			"code":    400,
			"message": "参数错误",
			"data":    nil,
		})
		return
	}
	username := strings.ToLower(stepAccountForm.Username)
	nickname := stepAccountForm.Nickname
	password := stepAccountForm.Password
	serial := stepAccountForm.Serial

	if !utils.CheckRegxp(username, constant.RegUsername) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户名格式错误",
			"data":    nil,
		})
		return
	}
	if !utils.CheckRegxp(password, constant.RegPassword) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码格式错误",
			"data":    nil,
		})
		return
	}

	result := MYSQL.First(&model.User{}, "username = ?", username)
	if result.Error == nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户名已存在",
			"data":    nil,
		})
		return
	}
	result = MYSQL.First(&model.RegFlow{}, "username = ?", username)
	if result.Error == nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户名已存在",
			"data":    nil,
		})
		return
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "什么动静",
			"detail":  err.Error(),
			"data":    nil,
		})
		return
	}

	uuidV4, err := uuid.NewV4()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "什么动静",
			"detail":  err.Error(),
			"data":    nil,
		})
		log.Printf("未能产生uuid：%v", err)
		return
	}
	state := uuidV4.String()

	postForm := map[string]interface{}{
		"step":     3,
		"username": username,
		"nickname": nickname,
		"password": string(pass),
		"ms_state": state,
	}

	MYSQL.Model(&model.RegFlow{}).Where("serial = ?", serial).Updates(postForm)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建成功",
		"data": map[string]interface{}{
			"url":         "/reg/flow/4",
			"link_start":  utils.GenerateLinkStart(state),
			"link_remake": utils.GenerateLinkRemake(),
		},
	})
}
