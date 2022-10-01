package api_reg

import (
	"net/http"
	"sso_gin/db"
	"sso_gin/model"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func HandleStepStart(ctx *gin.Context) {
	MYSQL := *db.MYSQL
	serial := uuid.NewV4()
	MYSQL.Create(&model.RegFlow{
		Serial: serial.String(),
	})
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "流程启动",
		"serial":  serial,
	})
}
