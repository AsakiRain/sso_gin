package api

import (
	"net/http"
	"sso_gin/constant"
	"sso_gin/db"
	"sso_gin/model"
	"sso_gin/utils"
	"unsafe"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const (
	RegUsername = constant.RegUsername
	RegPassword = constant.RegPassword
	RegEmail    = constant.RegEmail
)

func HandleRegister(ctx *gin.Context) {
	MYSQL := *db.MYSQL
	var registerForm model.RegisterForm
	err := ctx.ShouldBindJSON(&registerForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	name := registerForm.Username
	nickname := registerForm.Nickname
	password := registerForm.Password
	email := registerForm.Email
	code := registerForm.Code

	if !utils.CheckRegxp(name, RegUsername) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户名格式错误",
		})
		return
	}
	if !utils.CheckRegxp(password, RegPassword) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码格式错误",
		})
		return
	}
	if !utils.CheckRegxp(email, RegEmail) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "邮箱格式错误",
		})
		return
	}
	if code == "" {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "验证码（暂时）不能为空",
		})
		return
	}
	var user model.User
	result := MYSQL.First(&user, "username = ?", name)
	if result.Error == nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户名已存在",
		})
		return
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "什么动静",
			"detail":  err.Error(),
		})
		return
	}
	MYSQL.Create(&model.User{
		Username: name,
		Nickname: nickname,
		Password: *(*string)(unsafe.Pointer(&pass)),
		Email:    email,
	})
	userJwt := model.UserJwt{
		Username: name,
		Nickname: nickname,
		Email:    email,
		Role:     "user",
	}
	jwtToken, err := utils.GenerateToken(userJwt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "什么动静",
			"detail":  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "注册成功",
		"token":   jwtToken,
	})
}
