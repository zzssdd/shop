package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var SecretKey = []byte("as234sdgs")

type MyClaim struct {
	jwt.StandardClaims
	// 自定义的用户信息
	UserName string `json:"user_name"`
}

func NewToken(name string) (string, error) {
	user := MyClaim{
		jwt.StandardClaims{
			// 现在 + 加上传的过期时间
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
			Issuer:    "micro_gin_vue",
		},
		name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, user)
	return token.SignedString(SecretKey)
}

func AuthToken(tokenstring string) (*MyClaim, error) {
	token, err := jwt.ParseWithClaims(tokenstring, &MyClaim{}, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	claims, is_ok := token.Claims.(*MyClaim)
	if is_ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("token不合法")
}
