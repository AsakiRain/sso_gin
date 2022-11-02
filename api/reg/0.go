package api_reg

import (
	"log"
	"net/http"
	"sso_gin/db"
	"sso_gin/model"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func HandleStepStart(ctx *gin.Context) {
	MYSQL := *db.MYSQL

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
	serial := uuidV4.String()
	MYSQL.Create(&model.RegFlow{
		Serial: serial,
	})
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "流程启动",
		"data": map[string]interface{}{
			"serial": serial,
		},
	})
}
