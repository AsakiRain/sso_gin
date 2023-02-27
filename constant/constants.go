package constant

const (
	RegUsername = "^[a-zA-Z0-9_]{3,20}$"
	RegNickname = "^.{1,20}$"
	RegPassword = "^[\x21-\x7e]{8,36}$"
	RegEmail    = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)?$"
)
