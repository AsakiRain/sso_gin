package model

type EmailCaptcha struct {
	Email     string
	Code      string
	ExpiresAt int64
}
