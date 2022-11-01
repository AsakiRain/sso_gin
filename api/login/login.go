package api_login

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"sso_gin/db"
	"sso_gin/model"
	"sso_gin/utils"
)

func HandleLogin(ctx *gin.Context) {
	MYSQL := *db.MYSQL
	var loginForm model.LoginForm
	err := ctx.ShouldBindJSON(&loginForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}

	username := loginForm.Username
	password := loginForm.Password

	var user model.User
	result := MYSQL.First(&user, "username = ?", username)
	if result.Error != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "账号不存在",
		})
		return
	}
	nickname := user.Nickname
	email := user.Email
	role := user.Role
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码错误",
		})
		return
	}
	userJwt := model.UserJwt{
		Username: username,
		Nickname: nickname,
		Email:    email,
		Role:     role,
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
		"message": "登录成功",
		"token":   jwtToken,
	})
}
