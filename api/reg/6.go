package api_reg

import (
	"net/http"
	"sso_gin/db"
	"sso_gin/model"

	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"
)

func HandleStepPreference(ctx *gin.Context) {
	MYSQL := *db.MYSQL
	var serialForm model.SerialForm
	ctx.ShouldBindBodyWith(&serialForm, binding.JSON)
	// 这里也不用担心绑定失败，因为中间件试过了
	serial := serialForm.Serial

	//TODO

	postForm := map[string]interface{}{
		"step": 6,
	}
	MYSQL.Model(&model.RegFlow{}).Where("serial = ?", serial).Updates(postForm)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    20000,
		"message": "TODO",
		"data": map[string]interface{}{
			"url": "/reg/flow/7",
		},
	})
}
