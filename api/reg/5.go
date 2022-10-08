package api_reg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleStepQq(ctx *gin.Context) {
	// TODO
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "TODO",
		"url":     "/reg/flow/6",
	})
}
