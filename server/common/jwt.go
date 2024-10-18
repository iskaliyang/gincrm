package common

import (
	"gincrm/global"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Claims struct {
	Uid int64 `json:"uid"`
	jwt.RegisteredClaims
}

func GenToken(uid int64) (string, error) {
	var signingKey = []byte(global.Config.Jwt.SigningKey)
	var expiredTime = global.Config.Jwt.ExpiredTime

	claims := Claims{uid, jwt.RegisteredClaims{
		ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Duration(expiredTime) * time.Second)},
		Issuer:    "gin-crm",
	}}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(signingKey)
	return token, err
}
