package model

type EmailForm struct {
	Email string `json:"email"`
}
type EmailCaptcha struct {
	Email     string
	Code      string
	ExpiresAt int64
}
