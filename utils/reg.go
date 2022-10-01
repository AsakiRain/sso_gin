package utils

import "regexp"

func CheckRegxp(value string, reg string) bool {
	regxp := regexp.MustCompile(reg)
	return regxp.MatchString(value)
}
