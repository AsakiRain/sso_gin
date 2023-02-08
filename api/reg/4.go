package api_reg

import (
	"fmt"
	"log"
	"net/http"
	"sso_gin/db"
	"sso_gin/model"
	"sso_gin/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	uuid "github.com/satori/go.uuid"
)

func HandleStepMs(ctx *gin.Context) {
	MYSQL := *db.MYSQL
	var stepMsForm model.StepMsForm
	err := ctx.ShouldBindBodyWith(&stepMsForm, binding.JSON)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    40001,
			"message": "参数错误",
			"data":    nil,
		})
		return
	}
	msToken := stepMsForm.MsToken
	msState := stepMsForm.MsState
	serial := stepMsForm.Serial

	var regFlow model.RegFlow
	MYSQL.First(&regFlow, "serial = ?", serial)
	// 这里不用判断记录是否存在，因为中间件会检查的
	if regFlow.MsState == nil || *regFlow.MsState != msState {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    42210,
			"message": "state不匹配",
			"data":    nil,
		})
		return
	}
	clearForm := map[string]interface{}{
		"ms_step":                0,
		"ms_tip":                 nil,
		"ms_state":               nil,
		"ms_end":                 0,
		"minecraft_id":           nil,
		"minecraft_name":         nil,
		"minecraft_skins":        nil,
		"minecraft_capes":        nil,
		"minecraft_entitlements": nil,
	}
	// 预先重置整个过程，防止上次失败的记录被再次获取
	// 这里清空state是为了防止重复提交
	MYSQL.Model(&regFlow).Where("serial = ?", serial).Updates(clearForm)

	go utils.LinkStart(serial, msToken)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    20000,
		"message": "流程启动",
		"data":    nil,
	})
}

func HandleGenerateLink(ctx *gin.Context) {
	MYSQL := *db.MYSQL
	var serialForm model.SerialForm
	err := ctx.ShouldBindBodyWith(&serialForm, binding.JSON)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    40001,
			"message": "参数错误",
			"data":    nil,
		})
		return
	}
	serial := serialForm.Serial

	uuidV4, err := uuid.NewV4()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    50000,
			"message": fmt.Sprintf("state生成失败: %s", err.Error()),
			"data":    nil,
		})
		log.Printf("未能产生state：%v", err)
		return
	}
	state := uuidV4.String()

	postForm := map[string]interface{}{
		"ms_state": state,
	}

	MYSQL.Model(&model.RegFlow{}).Where("serial = ?", serial).Updates(postForm)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    20000,
		"message": "创建成功",
		"data": map[string]interface{}{
			"link_start":  utils.GenerateLinkStart(state),
			"link_remake": utils.GenerateLinkRemake(),
		},
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
	var msTip model.MsTip
	if regFlow.MsTip != nil {
		utils.ToStruct(&msTip, *regFlow.MsTip)
	}
	if regFlow.MsEnd == 1 {
		var minecraftSkins []model.MinecraftSkin
		var minecraftCapes []model.MinecraftCape
		var minecraftEntitlements model.MinecraftEntitlements
		utils.ToStruct(&minecraftSkins, *regFlow.MinecraftSkins)
		utils.ToStruct(&minecraftCapes, *regFlow.MinecraftCapes)
		utils.ToStruct(&minecraftEntitlements, *regFlow.MinecraftEntitlements)

		msMinecraft.MinecraftId = regFlow.MinecraftId
		msMinecraft.MinecraftName = regFlow.MinecraftName
		msMinecraft.MinecraftSkins = &minecraftSkins
		msMinecraft.MinecraftCapes = &minecraftCapes
		msMinecraft.MinecraftEntitlements = &minecraftEntitlements
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    20000,
		"message": "查询成功",
		"data": map[string]interface{}{
			"ms_step":   regFlow.MsStep,
			"ms_tip":    msTip,
			"ms_end":    regFlow.MsEnd,
			"minecraft": msMinecraft,
		},
	})
}
