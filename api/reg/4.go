package api_reg

import (
	"net/http"
	"sso_gin/db"
	"sso_gin/model"
	"sso_gin/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func HandleStepMs(ctx *gin.Context) {
	MYSQL := *db.MYSQL
	var stepMsForm model.StepMsForm
	err := ctx.ShouldBindBodyWith(&stepMsForm, binding.JSON)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}
	msToken := stepMsForm.MsToken
	msState := stepMsForm.MsState
	serial := stepMsForm.Serial

	var regFlow model.RegFlow
	MYSQL.Model(&regFlow).Where("serial = ?", serial).First(&regFlow)
	// 这里不用判断记录是否存在，因为中间件会检查的
	if regFlow.MsState != nil && *regFlow.MsState != msState {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "state不匹配",
		})
		return
	}
	clearForm := map[string]interface{}{
		"ms_step":  0,
		"ms_tip":   nil,
		"ms_state": nil,
		"ms_end":   0,
	}
	// 预先重置整个过程，防止上次失败的记录被再次获取
	// 这里清空state是为了防止重复提交
	MYSQL.Model(&regFlow).Where("serial = ?", serial).Updates(clearForm)

	go utils.LinkStart(serial, msToken)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "流程启动",
	})
}

func HandleMsQuery(ctx *gin.Context) {
	MYSQL := *db.MYSQL
	var serialForm model.SerialForm
	ctx.ShouldBindBodyWith(&serialForm, binding.JSON)
	// 这里也不用担心绑定失败，因为中间件试过了
	serial := serialForm.Serial

	var regFlow model.RegFlow
	MYSQL.Model(&regFlow).Where("serial = ?", serial).First(&regFlow)

	var msMinecraft model.MsMinecraft
	if regFlow.MsEnd == 1 {
		msMinecraft.MinecraftId = regFlow.MinecraftId
		msMinecraft.MinecraftName = regFlow.MinecraftName
		msMinecraft.MinecraftSkins = regFlow.MinecraftSkins
		msMinecraft.MinecraftCapes = regFlow.MinecraftCapes
		msMinecraft.MinecraftEntitlements = regFlow.MinecraftEntitlements
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":         200,
		"message":      "查询成功",
		"ms_step":      regFlow.MsStep,
		"ms_tip":       regFlow.MsTip,
		"ms_end":       regFlow.MsEnd,
		"ms_minecraft": msMinecraft,
	})
}
