package model

type RegFlow struct {
	Serial    string  `json:"serial" gorm:"not null"`
	Step      int     `json:"step" gorm:"default:0;not null"`
	Log       *string `json:"log"`
	MsStep    int     `json:"ms_step" gorm:"default:0;not null"`
	MsTip     *string `json:"ms_tip"`
	AcceptTos *bool   `json:"accept_tos" gorm:"default:false;not null"`
	Email     *string `json:"email"`
	Username  *string `json:"username"`
	Nickname  *string `json:"nickname"`
	Pass      *string `json:"password"`
	Salt      *string `json:"salt"`
	QqStep    int     `json:"qq_step" gorm:"default:0;not null"`
	QqTip     *string `json:"qq_tip"`
	CreatedAt int64   `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt int64   `json:"updated_at" gorm:"autoUpdateTime;not null"`
	DeletedAt *int64  `json:"deleted_at"`
}

type StepEmailForm struct {
	Email  string `json:"email" binding:"required"`
	Code   string `json:"code" binding:"required"`
	Serial string `json:"serial" binding:"required"`
}
