package utils

import (
	"sso_gin/config"
)

func GenerateLinkStart(state string) string {
	host := "login.microsoftonline.com"
	path := "consumers/oauth2/v2.0/authorize"
	query := map[string]interface{}{
		"client_id":     config.MsClientId,
		"response_type": "token",
		"redirect_uri":  config.MsRedirectUri,
		"scope":         "XboxLive.signin offline_access",
		"state":         state,
		"response_mode": "fragment",
	}
	link := GenerateLink(host, path, query)
	return link
}

func GenerateLinkRemake() string {
	return "https://www.ms.com"
}
