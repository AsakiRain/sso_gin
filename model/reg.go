package model

import (
	"time"

	"github.com/go-resty/resty/v2"
	"gorm.io/gorm"
)

type RegFlow struct {
	Serial                string         `json:"serial" gorm:"not null"`
	Step                  int            `json:"step" gorm:"default:0;not null"`
	Log                   *string        `json:"log"`
	MsStep                int            `json:"ms_step" gorm:"default:0;not null"`
	MsTip                 *string        `json:"ms_tip"`
	MsStatus              *string        `json:"ms_status" gorm:"default:'idle'"`
	MsState               *string        `json:"ms_state"`
	MinecraftId           *string        `json:"minecraft_id"`
	MinecraftName         *string        `json:"minecraft_name"`
	MinecraftSkins        *string        `json:"minecraft_skins"`
	MinecraftCapes        *string        `json:"minecraft_capes"`
	MinecraftEntitlements *string        `json:"minecraft_entitlements"`
	AcceptTos             *bool          `json:"accept_tos" gorm:"default:false;not null"`
	Email                 *string        `json:"email"`
	Username              *string        `json:"username"`
	Nickname              *string        `json:"nickname"`
	Password              *string        `json:"password"`
	Salt                  *string        `json:"salt"`
	QqNumber              *string        `json:"qq_number"`
	QqStatus              *string        `json:"qq_status" gorm:"default:'idle'"`
	CreatedAt             time.Time      `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt             time.Time      `json:"updated_at" gorm:"autoUpdateTime;not null"`
	DeletedAt             gorm.DeletedAt `json:"deleted_at"`
}

type SerialForm struct {
	Serial string `json:"serial" binding:"required"`
}

type StepTOSForm struct {
	AcceptTos *bool  `json:"accept_tos" binding:"required"`
	Serial    string `json:"serial" binding:"required"`
}

type StepEmailForm struct {
	Email  string `json:"email" binding:"required"`
	Code   string `json:"code" binding:"required"`
	Serial string `json:"serial" binding:"required"`
}

type StepAccountForm struct {
	Username string `json:"username" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Password string `json:"password" binding:"required"`
	Serial   string `json:"serial" binding:"required"`
}

type StepMsForm struct {
	MsToken string `json:"ms_token" binding:"required"`
	MsState string `json:"ms_state" binding:"required"`
	Serial  string `json:"serial" binding:"required"`
}

type MsTipForm struct {
	Info  string          `json:"info" binding:"required"`
	Error string          `json:"error" binding:"required"`
	Trace resty.TraceInfo `json:"trace" binding:"required"`
}

type MsXblForm struct {
	Properties   MsXblProperties `json:"Properties"`
	RelyingParty string          `json:"RelyingParty"`
	TokenType    string          `json:"TokenType"`
}

type MsXblProperties struct {
	AuthMethod string `json:"AuthMethod"`
	SiteName   string `json:"SiteName"`
	RpsTicket  string `json:"RpsTicket"`
}

type MsXblResp struct {
	DisplayClaims struct {
		Xui []struct {
			Uhs string `json:"uhs"`
		} `json:"xui"`
	} `json:"DisplayClaims"`
	Token        string `json:"Token"`
	NotAfter     string `json:"NotAfter"`
	IssueInstant string `json:"IssueInstant"`
}

type MsXblReturn struct {
	XblToken string `json:"xbl_token"`
	UserHash string `json:"user_hash"`
}

type MsXstsForm struct {
	Properties   MsXstsProperties `json:"Properties"`
	RelyingParty string           `json:"RelyingParty"`
	TokenType    string           `json:"TokenType"`
}

type MsXstsProperties struct {
	SandboxId  string `json:"SandboxId"`
	UserTokens []string
}

type MsXstsResp struct {
	Token         string `json:"Token"`
	NotAfter      string `json:"NotAfter"`
	IssueInstant  string `json:"IssueInstant"`
	DisplayClaims struct {
		Xui []struct {
			Uhs string `json:"uhs"`
		} `json:"xui"`
	} `json:"DisplayClaims"`
}

type MsXstsReturn struct {
	XstsToken string `json:"xsts_token"`
	UserHash  string `json:"user_hash"`
}

type MsMinecraftForm struct {
	IdentityToken string `json:"identityToken"`
}

type MsMinecraftResp struct {
	Username    string     `json:"username"`
	Roles       []struct{} `json:"roles"`
	AccessToken string     `json:"access_token"`
	TokenType   string     `json:"token_type"`
	ExpiresIn   int        `json:"expires_in"`
}

type MsMinecraftReturn struct {
	AccessToken string `json:"access_token"`
}

type MsEntitlementsForm struct{}

type MsEntitlementsResp struct {
	Items []struct {
		Name      string `json:"name"`
		Signature string `json:"signature"`
	} `json:"items"`
	Signature string `json:"signature"`
	KeyID     string `json:"keyId"`
}

type MsEntitlementsReturn MsEntitlementsResp

type MsProfileForm struct{}

type MsProfileResp struct {
	Id    string          `json:"id"`
	Name  string          `json:"name"`
	Skins []MinecraftSkin `json:"skins"`
	Capes []MinecraftCape `json:"capes"`
}

type MsProfileReturn MsProfileResp

type MsMinecraft struct {
	MinecraftId           *string                `json:"minecraft_id"`
	MinecraftName         *string                `json:"minecraft_name"`
	MinecraftSkins        *[]MinecraftSkin       `json:"minecraft_skins"`
	MinecraftCapes        *[]MinecraftCape       `json:"minecraft_capes"`
	MinecraftEntitlements *MinecraftEntitlements `json:"minecraft_entitlements"`
}

type MinecraftSkin struct {
	Id      string `json:"id"`
	State   string `json:"state"`
	Url     string `json:"url"`
	Variant string `json:"variant"`
	Alias   string `json:"alias"`
}
type MinecraftCape struct {
	Id    string `json:"id"`
	State string `json:"state"`
	Url   string `json:"url"`
	Alias string `json:"alias"`
}

type MinecraftEntitlements MsEntitlementsResp

type MsTip struct {
	Info  string      `json:"info"`
	Error string      `json:"error"`
	Trace MyTraceInfo `json:"trace"`
}
type MyTraceInfo struct {
	DNSLookup      time.Duration
	ConnTime       time.Duration
	TCPConnTime    time.Duration
	TLSHandshake   time.Duration
	ServerTime     time.Duration
	ResponseTime   time.Duration
	TotalTime      time.Duration
	IsConnReused   bool
	IsConnWasIdle  bool
	ConnIdleTime   time.Duration
	RequestAttempt int
	RemoteAddr     struct {
		IP   string
		Port int
		Zone string
	}
}

type QqCodeForm struct {
	Qq     string `json:"qq" binding:"required"`
	Serial string `json:"serial" binding:"required"`
}

type QqCaptcha struct {
	Serial    string
	Qq        string
	Code      string
	ExpiresAt int64
}

type QqCheckForm struct {
	RpcSecret string `json:"rpc_secret" binding:"required"`
	Qq        string `json:"qq" binding:"required"`
	Code      string `json:"code" binding:"required"`
}

type StepQqForm struct {
	Serial string `json:"serial" binding:"required"`
}
