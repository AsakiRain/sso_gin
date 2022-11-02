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
			"code": 401,
			"msg":  "token过期",
			"data": nil,
		})
		return
	}
	log.Println(claims)
	var user model.User
	result := MYSQL.First(&user, "username = ?", claims.UserJwt.Username)
	if result.Error != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "用户不存在",
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取用户信息成功",
		"data": user,
	})
}
