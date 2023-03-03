package utils

import (
	"encoding/json"
	"fmt"
	"sso_gin/config"
	"sso_gin/db"
	"sso_gin/model"

	"github.com/go-resty/resty/v2"
)

func GenerateLinkStart(state string) string {
	host := "login.microsoftonline.com"
	path := "consumers/oauth2/v2.0/authorize"
	query := map[string]interface{}{
		"client_id":     config.MsClientId,
		"response_type": "token",
		"redirect_uri":  config.MsRedirectUri,
		"scope":         "XboxLive.signin offline_access",
		"state":         state,
		"response_mode": "fragment",
	}
	link := GenerateLink(host, path, query)
	return link
}

func GenerateLinkRemake() string {
	// return "https://login.microsoftonline.com/logout.srf"
	return "https://login.live.com/logout.srf"
}

func JudgeResp(resp *resty.Response, err error, msTipForm *model.MsTipForm, info string) bool {
	msTipForm.Trace = resp.Request.TraceInfo()
	if err != nil {
		msTipForm.Info = info
		msTipForm.Error = err.Error()
		return false
	}
	if resp.StatusCode() != 200 {
		msTipForm.Info = info
		msTipForm.Error = fmt.Sprintf("StatusCode: %d", resp.StatusCode())
		return false
	}
	return true
}

func GailyPass(serial string, step int, msTipForm model.MsTipForm) bool {
	MYSQL := *db.MYSQL

	commitForm := map[string]interface{}{
		"MsStep": step,
		"MsTip":  ToJson(&msTipForm),
	}
	MYSQL.Model(&model.RegFlow{}).Where("serial = ?", serial).Updates(commitForm)
	return true
}

func SadlyDie(serial string, step int, msTipForm model.MsTipForm) bool {
	MYSQL := *db.MYSQL

	commitForm := map[string]interface{}{
		"MsStep":   step,
		"MsTip":    ToJson(&msTipForm),
		"MsStatus": "failed",
	}
	MYSQL.Model(&model.RegFlow{}).Where("serial = ?", serial).Updates(commitForm)
	return false
}

func LinkStart(serial string, msToken string) bool {
	MYSQL := *db.MYSQL

	MYSQL.Model(&model.RegFlow{}).Where("serial = ?", serial).Update("ms_step", 1)

	msXblReturn, msTipForm := MsStepXbl(serial, msToken)
	if msTipForm.Error != "" {
		return SadlyDie(serial, 2, msTipForm)
	}
	msXstsReturn, msTipForm := MsStepXsts(serial, msXblReturn)
	if msTipForm.Error != "" {
		return SadlyDie(serial, 3, msTipForm)
	}
	msMinecraftReturn, msTipForm := MsStepMinecraft(serial, msXstsReturn)
	if msTipForm.Error != "" {
		return SadlyDie(serial, 4, msTipForm)
	}
	msEntitlementsReturn, msTipForm := MsStepEntitlements(serial, msMinecraftReturn)
	if msTipForm.Error != "" {
		return SadlyDie(serial, 5, msTipForm)
	}
	msProfileReturn, msTipForm := MsStepProfile(serial, msMinecraftReturn)
	if msTipForm.Error != "" {
		return SadlyDie(serial, 6, msTipForm)
	}

	updateForm := map[string]interface{}{
		"MsStep":                7,
		"MsTip":                 ToJson(&msTipForm),
		"MsStatus":              "succeed",
		"MinecraftId":           msProfileReturn.Id,
		"MinecraftName":         msProfileReturn.Name,
		"MinecraftSkins":        ToJson(msProfileReturn.Skins),
		"MinecraftCapes":        ToJson(msProfileReturn.Capes),
		"MinecraftEntitlements": ToJson(msEntitlementsReturn),
	}
	MYSQL.Model(&model.RegFlow{}).Where("serial = ?", serial).Updates(updateForm)
	return true
}

func MsStepXbl(serial string, msToken string) (model.MsXblReturn, model.MsTipForm) {
	// /////////////////////////
	// Authenticate with XBL //
	// /////////////////////////
	var msTipForm model.MsTipForm
	client := resty.New()

	msXblForm := model.MsXblForm{
		Properties: model.MsXblProperties{
			AuthMethod: "RPS",
			SiteName:   "user.auth.xboxlive.com",
			RpsTicket:  "d=" + msToken,
		},
		RelyingParty: "http://auth.xboxlive.com",
		TokenType:    "JWT",
	}

	resp, err := client.R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetBody(msXblForm).
		Post("https://user.auth.xboxlive.com/user/authenticate")

	if JudgeResp(resp, err, &msTipForm, "请求XBL认证失败") == false {
		return model.MsXblReturn{}, msTipForm
	}

	msXblResp := model.MsXblResp{}
	err = json.Unmarshal(resp.Body(), &msXblResp)
	if err != nil {
		msTipForm.Info = "解析XBL认证失败"
		msTipForm.Error = err.Error()
		return model.MsXblReturn{}, msTipForm
	}

	msXblReturn := model.MsXblReturn{
		XblToken: msXblResp.Token,
		UserHash: msXblResp.DisplayClaims.Xui[0].Uhs,
	}
	msTipForm.Info = "请求XBL认证成功"
	msTipForm.Error = ""

	GailyPass(serial, 2, msTipForm)
	return msXblReturn, msTipForm
}

func MsStepXsts(serial string, msXblReturn model.MsXblReturn) (model.MsXstsReturn, model.MsTipForm) {
	////////////////////////////
	// Authenticate with XSTS //
	////////////////////////////
	var msTipForm model.MsTipForm
	client := resty.New()

	msXstsForm := model.MsXstsForm{
		Properties: model.MsXstsProperties{
			SandboxId: "RETAIL",
			UserTokens: []string{
				msXblReturn.XblToken,
			},
		},
		RelyingParty: "rp://api.minecraftservices.com/",
		TokenType:    "JWT",
	}

	resp, err := client.R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetBody(msXstsForm).
		Post("https://xsts.auth.xboxlive.com/xsts/authorize")

	if JudgeResp(resp, err, &msTipForm, "请求XSTS认证失败") == false {
		return model.MsXstsReturn{}, msTipForm
	}

	msXstsResp := model.MsXstsResp{}
	err = json.Unmarshal(resp.Body(), &msXstsResp)
	if err != nil {
		msTipForm.Info = "解析XSTS认证失败"
		msTipForm.Error = err.Error()
		return model.MsXstsReturn{}, msTipForm
	}

	msXstsReturn := model.MsXstsReturn{
		XstsToken: msXstsResp.Token,
		UserHash:  msXblReturn.UserHash,
	}
	msTipForm.Info = "请求XSTS认证成功"
	msTipForm.Error = ""

	GailyPass(serial, 3, msTipForm)
	return msXstsReturn, msTipForm
}

func MsStepMinecraft(serial string, msXstsReturn model.MsXstsReturn) (model.MsMinecraftReturn, model.MsTipForm) {
	/////////////////////////////////
	// Authenticate with Minecraft //
	/////////////////////////////////
	var msTipForm model.MsTipForm
	client := resty.New()

	msMinecraftForm := model.MsMinecraftForm{
		IdentityToken: fmt.Sprintf("XBL3.0 x=%s;%s", msXstsReturn.UserHash, msXstsReturn.XstsToken),
	}

	resp, err := client.R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetBody(msMinecraftForm).
		Post("https://api.minecraftservices.com/authentication/login_with_xbox")

	if JudgeResp(resp, err, &msTipForm, "请求Minecraft认证失败") == false {
		return model.MsMinecraftReturn{}, msTipForm
	}

	msMinecraftResp := model.MsMinecraftResp{}
	err = json.Unmarshal(resp.Body(), &msMinecraftResp)
	if err != nil {
		msTipForm.Info = "解析Minecraft认证失败"
		msTipForm.Error = err.Error()
		return model.MsMinecraftReturn{}, msTipForm
	}

	msMinecraftReturn := model.MsMinecraftReturn{
		AccessToken: msMinecraftResp.AccessToken,
	}
	msTipForm.Info = "请求Minecraft认证成功"
	msTipForm.Error = ""

	GailyPass(serial, 4, msTipForm)
	return msMinecraftReturn, msTipForm
}

func MsStepEntitlements(serial string, msMinecraftReturn model.MsMinecraftReturn) (model.MsEntitlementsReturn, model.MsTipForm) {
	/////////////////////////////
	// Checking Game Ownership //
	/////////////////////////////
	var msTipForm model.MsTipForm
	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		SetHeader("Authorization", "Bearer "+msMinecraftReturn.AccessToken).
		Get("https://api.minecraftservices.com/entitlements/mcstore")

	if JudgeResp(resp, err, &msTipForm, "请求Minecraft Entitlements失败") == false {
		return model.MsEntitlementsReturn{}, msTipForm
	}

	msEntitlementsResp := model.MsEntitlementsResp{}
	err = json.Unmarshal(resp.Body(), &msEntitlementsResp)
	if err != nil {
		msTipForm.Info = "解析Minecraft Entitlements失败"
		msTipForm.Error = err.Error()
		return model.MsEntitlementsReturn{}, msTipForm
	}

	msEntitlementsReturn := model.MsEntitlementsReturn(msEntitlementsResp)
	msTipForm.Info = "请求Minecraft Entitlements成功"
	msTipForm.Error = ""

	GailyPass(serial, 5, msTipForm)
	return msEntitlementsReturn, msTipForm
}

func MsStepProfile(serial string, msMinecraftReturn model.MsMinecraftReturn) (model.MsProfileReturn, model.MsTipForm) {
	/////////////////////
	// Get the profile //
	/////////////////////
	var msTipForm model.MsTipForm
	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		SetHeader("Authorization", "Bearer "+msMinecraftReturn.AccessToken).
		Get("https://api.minecraftservices.com/minecraft/profile")

	if JudgeResp(resp, err, &msTipForm, "请求Minecraft Profile失败") == false {
		return model.MsProfileReturn{}, msTipForm
	}

	msProfileResp := model.MsProfileResp{}
	err = json.Unmarshal(resp.Body(), &msProfileResp)
	if err != nil {
		msTipForm.Info = "解析Minecraft Profile失败"
		msTipForm.Error = err.Error()
		return model.MsProfileReturn{}, msTipForm
	}

	msProfileReturn := model.MsProfileReturn(msProfileResp)
	msTipForm.Info = "请求Minecraft Profile成功"
	msTipForm.Error = ""

	GailyPass(serial, 6, msTipForm)
	return msProfileReturn, msTipForm
}
