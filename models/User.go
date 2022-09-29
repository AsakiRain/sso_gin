package models

type User struct {
	Uid       int     `json:"uid" gorm:"not null;primaryKey;autoIncrement"`
	Username  string  `json:"username" gorm:"not null"`
	Nickname  string  `json:"nickname" gorm:"not null"`
	Pass      string  `json:"password" gorm:"not null"`
	Salt      string  `json:"salt" gorm:"not null"`
	Email     string  `json:"email" gorm:"not null"`
	Phone     *string `json:"phone"`
	Avatar    *string `json:"avatar"`
	Role      string  `json:"role" gorm:"default:user;not null"`
	CreatedAt int64   `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt int64   `json:"updated_at" gorm:"autoUpdateTime;not null"`
	DeletedAt *int64  `json:"deleted_at"`
}
