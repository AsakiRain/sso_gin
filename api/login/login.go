package api_login

import (
	"net/http"
	"sso_gin/db"
	"sso_gin/model"
	"sso_gin/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/bcrypt"
)

func HandleLogin(ctx *gin.Context) {
	MYSQL := *db.MYSQL
	var loginForm model.LoginForm
	err := ctx.ShouldBindBodyWith(&loginForm, binding.JSON)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
			"data":    nil,
		})
		return
	}

	username := loginForm.Username
	password := loginForm.Password

	var user model.User
	result := MYSQL.First(&user, "username = ?", username)
	if result.Error != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    422,
			"message": "账号不存在",
			"data":    nil,
		})
		return
	}
	nickname := user.Nickname
	email := user.Email
	role := user.Role
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    422,
			"message": "密码错误",
			"data":    nil,
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
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "登录成功",
		"data": map[string]interface{}{
			"token": jwtToken,
		},
	})
}
