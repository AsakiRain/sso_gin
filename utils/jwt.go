package utils

import (
	"sso_gin/config"
	"sso_gin/model"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte(config.JwtSecret)

func GenerateToken(UserJwt model.UserJwt) (string, error) {
	expiresTime := time.Now().Add(time.Hour * 24)

	claims := &model.MyClaims{
		UserJwt: UserJwt,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresTime),
			Issuer:    "sso_gin",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}

func ParseToken(tokenString string) (*model.MyClaims, error) {
	claims := &model.MyClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	// 若token只是过期claims是有数据的，若token无法解析claims无数据
	return claims, err
}
