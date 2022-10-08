package model

import "github.com/golang-jwt/jwt/v4"

type User struct {
	Uid       int     `json:"uid" gorm:"not null;primaryKey;autoIncrement"`
	Username  string  `json:"username" gorm:"not null"`
	Nickname  string  `json:"nickname" gorm:"not null"`
	Password  string  `json:"password" gorm:"not null"`
	Salt      string  `json:"salt" gorm:"not null"`
	Email     string  `json:"email" gorm:"not null"`
	Phone     *string `json:"phone"`
	Avatar    *string `json:"avatar"`
	Role      string  `json:"role" gorm:"default:user;not null"`
	CreatedAt int64   `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt int64   `json:"updated_at" gorm:"autoUpdateTime;not null"`
	DeletedAt *int64  `json:"deleted_at"`
}

type UserJwt struct {
	Username string
	Nickname string
	Role     string
}

type MyClaims struct {
	UserJwt UserJwt
	jwt.RegisteredClaims
}

type UserInfo struct {
	Uid       int    `json:"uid"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	DeletedAt int64  `json:"deleted_at"`
}
