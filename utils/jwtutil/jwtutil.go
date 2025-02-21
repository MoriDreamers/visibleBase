package jwtutil

import (
	"JWT-TEST/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtSignKey = []byte(config.JwtSignKey)

// 自定义声明类型
type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// 2.封装生成token的方法
func GenToken(username string) (string, error) {
	claims := MyCustomClaims{
		"bar",
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(config.JwtExpTime))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "Mori",
			Subject:   "Mori",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(JwtSignKey) //签名并返回token字符串
	return ss, err
}
