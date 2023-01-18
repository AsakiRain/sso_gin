package api_reg

import (
	"fmt"
	"net/http"
	"sso_gin/db"
	"sso_gin/model"
	"sso_gin/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func HandleStepDone(ctx *gin.Context) {
	MYSQL := *db.MYSQL

	var serialForm model.SerialForm
	ctx.ShouldBindBodyWith(&serialForm, binding.JSON)
	serial := serialForm.Serial

	var regFlow model.RegFlow
	MYSQL.First(&regFlow, "serial = ?", serial)
	//不用检查是否存在，除非中间件见鬼了
	username := *regFlow.Username
	nickname := *regFlow.Nickname
	password := *regFlow.Password
	email := *regFlow.Email
	salt := "NOT USED"
	role := "user"

	user := model.User{
		Username: username,
		Nickname: nickname,
		Password: password,
		Email:    email,
		Salt:     salt,
		Role:     role,
	}
	minecraft := model.Minecraft{
		Username:     username,
		Uuid:         *regFlow.MinecraftId,
		Name:         *regFlow.MinecraftName,
		Skins:        *regFlow.MinecraftSkins,
		Capes:        *regFlow.MinecraftCapes,
		Entitlements: *regFlow.MinecraftEntitlements,
	}
	MYSQL.Create(&user)
	MYSQL.Create(&minecraft)

	userJwt := model.UserJwt{
		Username: username,
		Nickname: nickname,
		Email:    email,
		Role:     role,
	}
	jwtToken, err := utils.GenerateToken(userJwt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    50000,
			"message": fmt.Sprintf("生成token失败: %v", err),
			"data":    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    20000,
		"message": "注册成功",
		"data": map[string]interface{}{
			"token": jwtToken,
		},
	})
}
