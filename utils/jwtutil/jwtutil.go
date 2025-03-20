package jwtutil

import (
	"errors"
	"time"
	"visibleBase/config"
	"visibleBase/utils/logs"

	"github.com/gin-gonic/gin"
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
		username,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(config.JwtExpTime))), //过期时间按分钟算，在配置文件里修改
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

func ParseToken(ss string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(ss, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSignKey, nil
	})
	if err != nil {
		logs.Error(gin.H{"err": nil, "Status": 401}, "解析token失败")
		return nil, err
	}
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		logs.Warning(gin.H{"err": nil, "Status": 401}, "token无效")
		return nil, errors.New("token无效")
	}

}
