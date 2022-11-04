package api_reg

import (
	"fmt"
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
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": fmt.Sprintf("uuid生成失败: %s", err.Error()),
			"data":    nil,
		})
		log.Printf("uuid生成失败：%v", err)
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
