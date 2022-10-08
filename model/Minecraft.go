package model

type Minecraft struct {
	Uid          int    `json:"uid" gorm:"not null;primaryKey;autoIncrement"`
	Username     string `json:"username" gorm:"not null"`
	Uuid         string `json:"uuid" gorm:"not null"`
	Name         string `json:"name" gorm:"not null"`
	Skins        string `json:"skins" gorm:"not null"`
	Capes        string `json:"capes" gorm:"not null"`
	Entitlements string `json:"entitlements" gorm:"not null"`
	CreatedAt    int64  `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt    int64  `json:"updated_at" gorm:"autoUpdateTime;not null"`
	DeletedAt    *int64 `json:"deleted_at"`
}
