package api_user

import (
	"log"
	"net/http"
	"sso_gin/db"
	"sso_gin/model"
	"sso_gin/utils"

	"github.com/gin-gonic/gin"
)

func HandleUserInfo(ctx *gin.Context) {
	MYSQL := *db.MYSQL
	auth := ctx.Request.Header.Get("Authorization")
	claims, err := utils.ParseToken(auth)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": 40103,
			"msg":  "token过期",
			"data": nil,
		})
		return
	}
	log.Println(claims)
	var userInfo model.UserInfo
	result := MYSQL.Model(&model.User{}).First(&userInfo, "username = ?", claims.UserJwt.Username)
	if result.Error != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": 42201,
			"msg":  "用户不存在",
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 20000,
		"msg":  "获取用户信息成功",
		"data": userInfo,
	})
}
