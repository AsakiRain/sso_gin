package utils

import "net/url"

func GenerateLink(host string, path string, query map[string]interface{}) string {
	u := url.URL{}
	u.Scheme = "https"
	u.Host = host
	u.Path = path
	q := u.Query()
	for k, v := range query {
		q.Add(k, v.(string))
	}
	u.RawQuery = q.Encode()
	return u.String()
}
