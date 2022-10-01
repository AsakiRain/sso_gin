package model

type EmailForm struct {
	Email string `json:"email" binding:"required"`
}
type EmailCaptcha struct {
	Email     string
	Code      string
	ExpiresAt int64
}
