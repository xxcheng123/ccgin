package jwtool

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type User struct {
	UserID               int    `json:"userId"`
	Username             string `json:"username"`
	Version              int    `json:"version"`
	jwt.RegisteredClaims        // v5版本新加的方法
}

type JWTool struct {
	secretKey string
}

func (j *JWTool) Generate(userId int, username string, version int, duration time.Duration) (string, error) {
	claims := User{
		UserID:   userId,
		Username: username,
		Version:  version,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)), // 过期时间24小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),               // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),               // 生效时间
		},
	}
	// 使用HS256签名算法
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := t.SignedString([]byte(j.secretKey))

	return s, err
}

func (j *JWTool) Parse(token string) (*User, error) {
	t, err := jwt.ParseWithClaims(token, &User{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := t.Claims.(*User)
	if !ok || !t.Valid {
		return nil, err
	}

	return claims, nil
}

func New(secretKey string) *JWTool {
	return &JWTool{secretKey}
}
