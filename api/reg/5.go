package api_reg

import (
	"fmt"
	"net/http"
	"sso_gin/config"
	"sso_gin/constant"
	"sso_gin/db"
	"sso_gin/model"
	"sso_gin/utils"
	"time"

	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"
)

func HandleStepQq(ctx *gin.Context) {
	MYSQL := *db.MYSQL
	var serialForm model.SerialForm
	ctx.ShouldBindBodyWith(&serialForm, binding.JSON)
	// 这里也不用担心绑定失败，因为中间件试过了
	serial := serialForm.Serial

	var regFlow model.RegFlow
	MYSQL.Model(&regFlow).Where("serial = ?", serial).First(&regFlow)

	if regFlow.QqStatus == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    42212,
			"message": "未开始QQ验证",
			"data":    nil,
		})
		return
	}

	if *regFlow.MsStatus == "waiting" {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    42211,
			"message": "还没有完成验证",
			"data":    nil,
		})
		return
	}

	postForm := map[string]interface{}{
		"step": 5,
	}
	MYSQL.Model(&model.RegFlow{}).Where("serial = ?", serial).Updates(postForm)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    20000,
		"message": "认证成功",
		"data": map[string]interface{}{
			"url": "/reg/flow/6",
		},
	})
}

func HandleGetQqCode(ctx *gin.Context) {
	MYSQL := *db.MYSQL
	CACHE := *db.CACHE
	var qqCodeForm model.QqCodeForm
	err := ctx.ShouldBindBodyWith(&qqCodeForm, binding.JSON)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    40001,
			"message": "参数错误",
			"data":    nil,
		})
		return
	}

	serial := qqCodeForm.Serial
	qq := qqCodeForm.Qq

	if !utils.CheckRegxp(qq, constant.RegQq) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    40002,
			"message": "QQ格式错误",
			"data":    nil,
		})
		return
	}

	cdKey := fmt.Sprintf("qq_cd_%s", qq)
	x, found := CACHE.Get(cdKey)
	if found {
		cdAt := x.(int64)
		if cdAt > time.Now().Unix() {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    42207,
				"message": fmt.Sprintf("操作频繁，请在%d秒后重试", cdAt-time.Now().Unix()),
				"data":    nil,
			})
			return
		}
	}

	code := utils.RandomCode(16)
	cacheKey := fmt.Sprintf("qq_captcha_%s", code)
	qqCaptcha := model.QqCaptcha{
		Serial:    serial,
		Qq:        qq,
		Code:      code,
		ExpiresAt: time.Now().Add(time.Minute * 10).UnixMilli(),
	}
	CACHE.Set(cacheKey, qqCaptcha, time.Minute*10)
	CACHE.Set(cdKey, time.Now().Add(time.Second*60).Unix(), time.Second*60)

	postForm := map[string]interface{}{
		"qq_number": qq,
		"qq_status": "waiting",
	}
	MYSQL.Model(&model.RegFlow{}).Where("serial = ?", serial).Updates(postForm)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    20000,
		"message": "生成验证码成功",
		"data": map[string]interface{}{
			"code": code,
		},
	})
}

func HandleCheckQqCode(ctx *gin.Context) {
	MYSQL := *db.MYSQL
	CACHE := *db.CACHE
	var qqCheckForm model.QqCheckForm
	err := ctx.ShouldBindJSON(&qqCheckForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    40001,
			"message": "参数错误",
			"data":    nil,
		})
		return
	}

	var serial string
	rpc_secret := qqCheckForm.RpcSecret
	qq := qqCheckForm.Qq
	code := qqCheckForm.Code

	if rpc_secret != config.QqRpcSecret {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    40002,
			"message": "RPC密钥错误",
			"data":    nil,
		})
		return
	}

	cacheKey := fmt.Sprintf("qq_captcha_%s", code)
	pass := false
	x, found := CACHE.Get(cacheKey)
	if found {
		qqCaptcha := x.(model.QqCaptcha)
		if qqCaptcha.Qq == qq && qqCaptcha.ExpiresAt > time.Now().Unix() {
			serial = qqCaptcha.Serial
			pass = true
		}
		CACHE.Delete(cacheKey)
	}
	if !pass {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    42206,
			"message": "验证码错误喵，主人",
			"data":    nil,
		})
		return
	}

	updateForm := map[string]interface{}{
		"step":      5,
		"qq_number": qq,
		"qq_status": "succeed",
	}
	MYSQL.Model(&model.RegFlow{}).Where("serial = ?", serial).Updates(updateForm)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    20000,
		"message": "验证成功喵，谢谢主人",
		"data":    nil,
	})
}
