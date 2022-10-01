package api_user

import (
	"log"
	"net/http"
	"sso_gin/db"
	"sso_gin/model"
	"sso_gin/utils"

	"github.com/gin-gonic/gin"
)

func HandleUserInfo(c *gin.Context) {
	MYSQL := *db.MYSQL
	auth := c.Request.Header.Get("Authorization")
	claims, err := utils.ParseToken(auth)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "token过期",
		})
		return
	}
	log.Println(claims)
	var user model.User
	result := MYSQL.First(&user, "username = ?", claims.UserJwt.Username)
	if result.Error != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "用户不存在",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取用户信息成功",
		"data": user,
	})
}
