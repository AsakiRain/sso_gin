package api_reg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleStepPerference(ctx *gin.Context) {
	// TODO
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "TODO",
	})
}
