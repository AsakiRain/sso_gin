package model

import (
	"time"

	"gorm.io/gorm"
)

type Minecraft struct {
	Uid          int            `json:"uid" gorm:"not null;primaryKey;autoIncrement"`
	Username     string         `json:"username" gorm:"not null"`
	Uuid         string         `json:"uuid" gorm:"not null"`
	Name         string         `json:"name" gorm:"not null"`
	Skins        string         `json:"skins" gorm:"not null"`
	Capes        string         `json:"capes" gorm:"not null"`
	Entitlements string         `json:"entitlements" gorm:"not null"`
	CreatedAt    time.Time      `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt    time.Time      `json:"updated_at" gorm:"autoUpdateTime;not null"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}
